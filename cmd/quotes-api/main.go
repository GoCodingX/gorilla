package main

import (
	"fmt"
	"log/slog"

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
		logger.Fatal("failed to read env file", err)

		return
	}

	// load config
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		logger.Fatal("failed to load config", err)

		return
	}

	// initialize router
	srv, err := handlers.NewRouter(&cfg)
	if err != nil {
		logger.Fatal("failed to initialize router", err)

		return
	}

	// start the http server
	logger.Info("starting the server", slog.String("port", cfg.Port))

	if err := srv.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		logger.Fatal("failed to start http server", err)
	}
}
