package main

import (
	"errors"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/GoCodingX/gorilla/internal/logger"
	"github.com/GoCodingX/gorilla/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	// load env vars from .env files
	err := godotenv.Load()
	if err != nil {
		var pathError *fs.PathError
		if !errors.As(err, &pathError) {
			logger.ErrorAndExit("failed to read env file", err)

			return
		}
	}

	// load config
	cfg, err := config.Load()
	if err != nil {
		logger.ErrorAndExit("failed to load config", err)

		return
	}

	// parse the apple certificate
	cert, err := utils.ParsePemCert(cfg.AppleRootCertificate)
	if err != nil {
		logger.ErrorAndExit("failed to process apple certificate", err)

		return
	}

	// initialize router
	srv := handlers.NewRouter(cfg, cert)

	logger.Info("starting the server", slog.String("port", cfg.Port))

	if err := http.ListenAndServe(cfg.Port, srv); err != nil {
		logger.ErrorAndExit("failed to start server", err)
	}
}
