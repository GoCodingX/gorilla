package handlers

import (
	"github.com/GoCodingX/gorilla/internal/clients/featureflag"
	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// initialize dependencies
	featureFlagClient := featureflag.NewClient(cfg.FeatureFlagAPIURL)

	// initialize service
	service := NewPaymentsService(&NewPaymentsServiceParams{
		FeatureFlagClient: featureFlagClient,
	})

	// routes
	r.Post("/v1/notifications", service.HandleNotification)

	return r
}
