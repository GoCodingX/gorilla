package handlers

import (
	"net/http"

	"github.com/GoCodingX/gorilla/internal/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) PostQuotes(c echo.Context) error {
	payload := new(api.CreateQuoteRequest)
	if err := c.Bind(payload); err != nil {
		return err
	}

	response := api.PostQuotes200JSONResponse{
		AuthorId: uuid.New(),
		Text:     "some text here",
	}

	return c.JSON(http.StatusCreated, response)
}
