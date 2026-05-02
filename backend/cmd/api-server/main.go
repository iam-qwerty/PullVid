package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pullvid/backend/internal/api"
	"pullvid/backend/internal/downloader"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	service := downloader.NewService("yt-dlp")
	if err := service.CheckDependencies(); err != nil {
		logger.Error("dependency check failed", "err", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	handler := api.NewHandler(logger, service)

	mux.HandleFunc("POST /video-info", handler.VideoInfo)
	mux.HandleFunc("GET /download", handler.Download)

	allowedOrigin := envOrDefault("CORS_ALLOWED_ORIGIN", "*")

	server := &http.Server{
		Addr:              ":" + envOrDefault("PORT", "8080"),
		Handler:           api.CORS(allowedOrigin)(mux),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("server shutdown failed", "err", err)
		}
	}()

	logger.Info("starting server", "addr", server.Addr, "cors_allowed_origin", allowedOrigin)

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server failed", "err", err)
		os.Exit(1)
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
