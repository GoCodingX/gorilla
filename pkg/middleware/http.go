package middleware

import (
	"time"

	"github.com/GoCodingX/gorilla/pkg/errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
)

const timeoutInSeconds = 10

// todo: have the timeout return json

// TimeoutMiddleware sets a request timeout middleware with a custom
// "request timed out" message and a configurable timeout duration.
var TimeoutMiddleware = middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	ErrorMessage: "request timed out",
	Timeout:      timeoutInSeconds * time.Second,
})

// OApiValidatorMiddleware creates an Echo middleware that validates requests
// against the provided OpenAPI spec and returns structured errors on validation failure.
func OApiValidatorMiddleware(swagger *openapi3.T) echo.MiddlewareFunc {
	return oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapimiddleware.Options{
		ErrorHandler:      errors.OApiErrorHandler(),
		MultiErrorHandler: errors.MultiErrorHandler(),
		Options: openapi3filter.Options{
			MultiError: true,
		},
	})
}
