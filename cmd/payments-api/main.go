package main

import (
	"log/slog"
	"net/http"

	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/GoCodingX/gorilla/internal/logger"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func main() {
	// load env vars from .env files
	err := godotenv.Load()
	if err != nil {
		logger.ErrorAndExit("failed to read env file", err)

		return
	}

	// load config
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		logger.ErrorAndExit("failed to load config", err)

		return
	}

	// initialize router
	srv := handlers.NewRouter(&cfg)

	logger.Info("starting the server", slog.String("port", cfg.Port))

	if err := http.ListenAndServe(cfg.Port, srv); err != nil {
		logger.ErrorAndExit("failed to start server", err)
	}
}
