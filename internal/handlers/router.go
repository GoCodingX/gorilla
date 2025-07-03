package handlers

import (
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	middleware2 "github.com/GoCodingX/gorilla/pkg/middleware"
	pkgmiddleware "github.com/GoCodingX/gorilla/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(service *QuotesService, swagger *openapi3.T) (*echo.Echo, error) {
	e := echo.New()

	// Hide the default console output from echo
	e.HideBanner = true
	e.HidePort = true

	// middlewares
	e.Use(middleware.RequestID())
	e.Use(basicAuthMiddleware)
	e.Use(middleware2.OApiValidatorMiddleware(swagger))
	e.Use(middleware2.TimeoutMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = pkgmiddleware.CustomHTTPErrorHandler

	// register routes
	openapi.RegisterHandlers(e, service)

	return e, nil
}
