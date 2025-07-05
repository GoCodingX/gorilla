package handlers_test

import (
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

// createTestEchoContext returns a minimal echo.Context for testing
func createTestEchoContext() echo.Context {
	e := echo.New()
	req := httptest.NewRequest("", "/", nil)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}
