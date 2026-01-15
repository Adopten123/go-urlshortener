package main

import (
	"go-urlshortener/internal/config"
	"go-urlshortener/internal/lib/logger/sl"
	"go-urlshortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

	logger := setupLogger(config.Env)
	logger.Info("starting url-shortener...", slog.String("env", config.Env))
	logger.Debug("debug messages are enabled...")

	storage, err := sqlite.NewStorage(config.StoragePath)
	if err != nil {
		logger.Error("failed to init sqlite storage", sl.Err(err))
		os.Exit(1)
	}

	id, err := storage.SaveURL("https://github.com/", "github")
	if err != nil {
		logger.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}
	logger.Info("saved url", slog.Int64("id", id))

	id, err = storage.SaveURL("https://github.com/", "github")
	if err != nil {
		logger.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}
	logger.Info("saved url", slog.Int64("id", id))

	_ = storage
	// TODO: init router. Libs: chi, "chi render"
	// TODO: run server.
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
