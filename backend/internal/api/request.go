package api

import (
	"errors"
	"net/http"

	"pullvid/backend/internal/downloader"
)

var (
	errMissingURL     = errors.New("url query param is required")
	errInvalidFormat  = errors.New("format must be one of: mp4, webm")
	errInvalidQuality = errors.New("quality must be one of: 360, 480, 720, 1080")
	allowedFormats    = map[string]struct{}{"mp4": {}, "webm": {}}
	allowedQualities  = map[string]struct{}{"360": {}, "480": {}, "720": {}, "1080": {}}
)

func parseDownloadSpec(r *http.Request) (downloader.DownloadSpec, error) {
	spec := downloader.DownloadSpec{
		URL:     r.URL.Query().Get("url"),
		Format:  r.URL.Query().Get("format"),
		Quality: r.URL.Query().Get("quality"),
	}

	if spec.URL == "" {
		return downloader.DownloadSpec{}, errMissingURL
	}

	if spec.Format == "" {
		spec.Format = "mp4"
	}

	if spec.Quality == "" {
		spec.Quality = "720"
	}

	if _, ok := allowedFormats[spec.Format]; !ok {
		return downloader.DownloadSpec{}, errInvalidFormat
	}

	if _, ok := allowedQualities[spec.Quality]; !ok {
		return downloader.DownloadSpec{}, errInvalidQuality
	}

	return spec, nil
}
