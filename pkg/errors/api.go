package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ConvertEchoToApiError extracts and converts the payload inside an echo.HTTPError
// into a generated.ErrorResponse for consistent structured error handling across the API.
// Returns an error if the payload is not of type generated.ErrorResponse.
func ConvertEchoToApiError(err *echo.HTTPError) (*openapi.ErrorResponse, error) {
	errorResponse, ok := err.Message.(*openapi.ErrorResponse)
	if ok {
		return errorResponse, nil
	}

	msg, ok := err.Message.(string)
	if !ok {
		return nil, fmt.Errorf(
			"convertEchoToApiError: expected err.Message to be string, got %T (%v)",
			err.Message,
			err.Message,
		)
	}

	return &openapi.ErrorResponse{
		Code:    err.Code,
		Status:  http.StatusText(err.Code),
		Message: msg,
	}, nil
}

// OApiErrorHandler returns an Echo HTTP error handler that converts
// *echo.HTTPError instances produced by request validation failures into
// structured API error responses.
// Additionally, this ensures validation errors are returned consistently in generated.ErrorResponse
// format, allowing clients to receive detailed field-level validation feedback.
func OApiErrorHandler() func(c echo.Context, err *echo.HTTPError) error {
	return func(c echo.Context, err *echo.HTTPError) error {
		responsePayload, conversionErr := ConvertEchoToApiError(err)
		if conversionErr != nil {
			return fmt.Errorf("failed to conver to api error: %w", conversionErr)
		}

		if bindErr := c.Bind(responsePayload); bindErr != nil {
			return fmt.Errorf("failed to bind error response body: %w", bindErr)
		}

		return echo.NewHTTPError(err.Code, responsePayload)
	}
}

// MultiErrorHandler returns a callback that converts an openapi3.MultiError
// into a structured *echo.HTTPError with a generated.ErrorResponse payload.
// It extracts validation error details from the MultiError and formats them
// into a consistent API error response with HTTP 400 Bad Request.
func MultiErrorHandler() func(multiError openapi3.MultiError) *echo.HTTPError {
	return func(multiError openapi3.MultiError) *echo.HTTPError {
		status := http.StatusBadRequest

		response := &openapi.ErrorResponse{
			Code:    status,
			Message: "request validation failed",
			Status:  http.StatusText(status),
		}

		var details []openapi.Detail

		for _, me := range multiError {
			var schemaErr *openapi3.SchemaError

			if !errors.As(me, &schemaErr) {
				continue
			}

			var schemaMultiErr openapi3.MultiError
			if errors.As(schemaErr.Origin, &schemaMultiErr) {
				for _, sme := range schemaMultiErr {
					if errors.As(sme, &schemaErr) {
						details = append(details, openapi.Detail{
							Field:   strings.Join(schemaErr.JSONPointer(), "."),
							Message: schemaErr.Reason,
						})
					}
				}
			}
		}

		response.Details = &details

		return echo.NewHTTPError(response.Code, response)
	}
}
