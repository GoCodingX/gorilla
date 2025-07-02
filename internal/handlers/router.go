package handlers

import (
	"fmt"

	"github.com/GoCodingX/gorilla/internal/config"
	"github.com/GoCodingX/gorilla/internal/repository/pg"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(cfg *config.Config, repo *pg.Repository) (*echo.Echo, error) {
	e := echo.New()

	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error loading swagger spec: %w", err)
	}

	// middlewares
	e.Use(middleware.RequestID())
	e.Use(oApiValidatorMiddleware(swagger))
	e.Use(timeoutMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// initialize service
	service := NewQuotesService(&NewQuotesServiceParams{
		Repo: repo,
	})

	// register routes
	openapi.RegisterHandlers(e, service)

	return e, nil
}
