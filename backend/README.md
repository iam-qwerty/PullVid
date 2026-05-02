# PullVid Backend

This app serves the Go API for the PullVid frontend.

## Layout

- `cmd/api-server`: application entrypoint
- `internal/api`: HTTP handlers and middleware
- `internal/downloader`: `yt-dlp` integration

## Run locally

```powershell
go run ./cmd/api-server
```

The server expects these binaries to be available in `PATH`:

- `yt-dlp`
- `ffmpeg`

## Environment variables

- `PORT`: server port, defaults to `8080`
- `CORS_ALLOWED_ORIGIN`: frontend origin for CORS, defaults to `*`
