package handlers

import (
	"net/http"

	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) GetQuotes(c echo.Context) error {
	quotes := []openapi.PostQuotes200JSONResponse{{}}

	return c.JSON(http.StatusOK, quotes)
}
