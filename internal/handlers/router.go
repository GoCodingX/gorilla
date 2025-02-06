package handlers

import (
	"crypto/x509"

	"github.com/GoCodingX/gorilla/internal/clients/featureflag"
	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/GoCodingX/gorilla/internal/jwtprocessor/apple"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg *config.Config, appleCert *x509.Certificate) *chi.Mux {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// initialize dependencies
	featureFlagClient := featureflag.NewClient(cfg.FeatureFlagAPIURL)
	appleJwtProcessor := apple.NewJwtProcessor(appleCert)

	// initialize service
	service := NewPaymentsService(&NewPaymentsServiceParams{
		FeatureFlagClient: featureFlagClient,
		JwtProcessor:      appleJwtProcessor,
	})

	// routes
	r.Post("/v1/apple-notifications", service.HandleAppleNotification)

	return r
}
