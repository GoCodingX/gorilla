package handlers

import (
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
)

// todo: add description here
// todo: make the timeout return json
var timeoutMiddleware = middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	ErrorMessage: "request timed out",
	Timeout:      10 * time.Second,
})

// oApiValidatorMiddleware creates an Echo middleware that validates requests
// against the provided OpenAPI spec and returns structured errors on validation failure.
func oApiValidatorMiddleware(swagger *openapi3.T) echo.MiddlewareFunc {
	return oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapimiddleware.Options{
		ErrorHandler:      oApiErrorHandler(),
		MultiErrorHandler: multiErrorHandler(),
		Options: openapi3filter.Options{
			MultiError: true,
		},
	})
}
