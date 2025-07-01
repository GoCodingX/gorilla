package handlers

import (
	"fmt"

	"github.com/GoCodingX/gorilla/internal/api"
	"github.com/GoCodingX/gorilla/internal/clients/featureflag"
	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(cfg *config.Config) (*echo.Echo, error) {
	e := echo.New()

	swagger, err := api.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error loading swagger spec: %w", err)
	}

	// middlewares
	e.Use(middleware.RequestID())
	e.Use(oApiValidatorMiddleware(swagger))
	e.Use(timeoutMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// initialize dependencies
	featureFlagClient := featureflag.NewClient(cfg.FeatureFlagAPIURL)

	// initialize service
	service := NewQuotesService(&NewQuotesServiceParams{
		FeatureFlagClient: featureFlagClient,
	})

	// register routes
	api.RegisterHandlers(e, service)

	return e, nil
}
