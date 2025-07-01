package handlers

import (
	"net/http"

	"github.com/GoCodingX/gorilla/internal/api"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) GetQuotes(c echo.Context) error {
	quotes := []api.PostQuotes200JSONResponse{{}}

	return c.JSON(http.StatusOK, quotes)
}
