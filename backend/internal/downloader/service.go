package downloader

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	binary string
}

type VideoInfo struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Duration  int    `json:"duration,omitempty"`
	Uploader  string `json:"uploader,omitempty"`
}

type DownloadSpec struct {
	URL     string
	Format  string
	Quality string
}

type Progress struct {
	Downloaded int64
	Total      int64
	Speed      string
	ETA        int
	Percent    string
}

type DownloadStream struct {
	Reader      io.ReadCloser
	Filename    string
	ContentType string
	Progress    <-chan Progress
	wait        func() error
}

func NewService(binary string) *Service {
	return &Service{binary: binary}
}

func (s *Service) CheckDependencies() error {
	for _, name := range []string{s.binary, "ffmpeg"} {
		if _, err := exec.LookPath(name); err != nil {
			return fmt.Errorf("%s is not installed or not in PATH: %w", name, err)
		}
	}

	return nil
}

func (s *Service) Probe(ctx context.Context, rawURL string) (VideoInfo, error) {
	probeCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		probeCtx,
		s.binary,
		"--dump-single-json",
		"--no-playlist",
		"--skip-download",
		rawURL,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return VideoInfo{}, commandError("failed to fetch video info", err, stderr.String())
	}

	var payload struct {
		Title     string `json:"title"`
		Thumbnail string `json:"thumbnail"`
		Duration  int    `json:"duration"`
		Uploader  string `json:"uploader"`
	}

	if err := json.Unmarshal(output, &payload); err != nil {
		return VideoInfo{}, fmt.Errorf("failed to parse video info: %w", err)
	}

	return VideoInfo{
		Title:     payload.Title,
		Thumbnail: payload.Thumbnail,
		Duration:  payload.Duration,
		Uploader:  payload.Uploader,
	}, nil
}

func (s *Service) PrepareDownload(ctx context.Context, spec DownloadSpec) (*DownloadStream, error) {
	formatSelector := fmt.Sprintf(
		"best[height<=%s][ext=%s]/best[height<=%s]",
		spec.Quality, spec.Format, spec.Quality,
	)

	filename, err := s.resolveFilename(ctx, formatSelector, spec.Format, spec.URL)
	if err != nil {
		filename = deriveFilename(spec.URL, spec.Format)
	}

	ext := strings.TrimPrefix(filepath.Ext(filename), ".")
	if ext == "" {
		ext = spec.Format
	}

	progressTemplate := "downloaded:%(progress.downloaded_bytes)d|total:%(progress.total_bytes_estimate)d|" +
		"speed:%(progress._speed_str)s|eta:%(progress.eta)d|pct:%(progress._percent_str)s"

	cmd := exec.CommandContext(ctx, s.binary,
		"-f", formatSelector,
		"--no-playlist",
		"--progress-template", progressTemplate,
		"--newline",
		"-o", "-",
		spec.URL,
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start yt-dlp: %w", err)
	}

	progressCh := make(chan Progress, 64)

	go func() {
		defer close(progressCh)
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			if p, ok := parseProgress(scanner.Text()); ok {
				select {
				case progressCh <- p:
				default:
				}
			}
		}
	}()

	return &DownloadStream{
		Reader:      stdout,
		Filename:    filename,
		ContentType: contentType(ext),
		Progress:    progressCh,
		wait:        cmd.Wait,
	}, nil
}

func (s *Service) resolveFilename(ctx context.Context, formatSelector, fallbackFormat, rawURL string) (string, error) {
	getCmd := exec.CommandContext(ctx, s.binary,
		"-f", formatSelector,
		"--no-playlist",
		"--get-filename",
		"-o", "%(title)s.%(ext)s",
		rawURL,
	)

	output, err := getCmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to resolve filename: %w", err)
	}

	filename := strings.TrimSpace(string(output))
	if filename == "" || filename == "-" {
		return "", fmt.Errorf("empty filename from yt-dlp")
	}

	return filename, nil
}

func (s *DownloadStream) Wait() error {
	return s.wait()
}

func (s *DownloadStream) Close() error {
	return s.Reader.Close()
}

func parseProgress(line string) (Progress, bool) {
	if !strings.HasPrefix(line, "downloaded:") {
		return Progress{}, false
	}

	var p Progress
	for _, part := range strings.Split(line, "|") {
		kv := strings.SplitN(part, ":", 2)
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "downloaded":
			p.Downloaded, _ = strconv.ParseInt(kv[1], 10, 64)
		case "total":
			p.Total, _ = strconv.ParseInt(kv[1], 10, 64)
		case "speed":
			p.Speed = kv[1]
		case "eta":
			p.ETA, _ = strconv.Atoi(kv[1])
		case "pct":
			p.Percent = kv[1]
		}
	}

	if p.Downloaded == 0 && p.Percent == "" {
		return Progress{}, false
	}

	return p, true
}

func contentType(ext string) string {
	switch ext {
	case "mp4":
		return "video/mp4"
	case "webm":
		return "video/webm"
	default:
		return "application/octet-stream"
	}
}

func deriveFilename(rawURL, format string) string {
	trimmed := strings.TrimRight(rawURL, "/")
	parts := strings.Split(trimmed, "/")
	name := strings.Split(parts[len(parts)-1], "?")[0]

	if name == "" || !strings.Contains(name, ".") {
		name = fmt.Sprintf("video_%d", time.Now().Unix())
	}

	return fmt.Sprintf("%s.%s", name, format)
}

func commandError(prefix string, err error, stderr string) error {
	stderr = strings.TrimSpace(stderr)
	if stderr == "" {
		return fmt.Errorf("%s: %w", prefix, err)
	}

	return fmt.Errorf("%s: %s", prefix, stderr)
}

