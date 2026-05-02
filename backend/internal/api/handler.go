package api

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"pullvid/backend/internal/downloader"
)

type downloadService interface {
	Probe(ctx context.Context, rawURL string) (downloader.VideoInfo, error)
	PrepareDownload(ctx context.Context, spec downloader.DownloadSpec) (*downloader.DownloadStream, error)
}

type Handler struct {
	logger  *slog.Logger
	service downloadService
}

func NewHandler(logger *slog.Logger, service downloadService) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) VideoInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req videoInfoRequest
	if err := decodeJSON(r, &req); err != nil {
		WriteJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid JSON body"})
		return
	}

	if req.URL == "" {
		WriteJSON(w, http.StatusBadRequest, errorResponse{Error: "url is required"})
		return
	}

	info, err := h.service.Probe(r.Context(), req.URL)
	if err != nil {
		h.logger.Warn("failed to fetch video info", "err", err)
		WriteJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, info)
}

func (h *Handler) Download(w http.ResponseWriter, r *http.Request) {
	spec, err := parseDownloadSpec(r)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	stream, err := h.service.PrepareDownload(r.Context(), spec)
	if err != nil {
		h.logger.Warn("failed to prepare download", "err", err)
		WriteJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	defer stream.Close()

	go h.logProgress(r.Context(), stream.Progress)

	w.Header().Set("Content-Disposition", `attachment; filename="`+stream.Filename+`"`)
	w.Header().Set("Content-Type", stream.ContentType)

	written, copyErr := io.Copy(w, stream.Reader)
	waitErr := stream.Wait()

	if copyErr != nil {
		h.logger.Warn("stream copy failed", "bytes_written", written, "err", copyErr)
		return
	}

	if waitErr != nil {
		if errors.Is(r.Context().Err(), context.Canceled) {
			h.logger.Info("client canceled download")
			return
		}

		h.logger.Warn("download process failed", "bytes_written", written, "err", waitErr)
	}
}

func (h *Handler) logProgress(ctx context.Context, ch <-chan downloader.Progress) {
	for {
		select {
		case p, ok := <-ch:
			if !ok {
				return
			}
			h.logger.Debug("download progress", "pct", p.Percent, "speed", p.Speed, "eta_s", p.ETA)
		case <-ctx.Done():
			return
		}
	}
}
