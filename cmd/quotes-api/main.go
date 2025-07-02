package main

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/GoCodingX/gorilla/internal/repository/pg"
	"github.com/GoCodingX/gorilla/pkg/logger"
	"github.com/GoCodingX/gorilla/pkg/migrate"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

	repo, err := initializeRepository(&cfg, err)
	if err != nil {
		logger.Fatal("failed to initialize repository", err)

		return
	}

	logger.Info("successfully initialized repository")

	// initialize router
	srv, err := handlers.NewRouter(&cfg, repo)
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

func initializeRepository(cfg *config.Config, err error) (*pg.Repository, error) {
	dbConn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DatabaseUrl)))

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed connect to db: %w", err)
	}

	dbClient := bun.NewDB(dbConn, pgdialect.New())

	err = migrate.Up(cfg.MigrationsDir, cfg.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	repo := pg.NewRepository(dbClient)

	return repo, nil
}
