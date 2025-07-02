package handlers

import (
	"net/http"

	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) PostQuotes(c echo.Context) error {
	payload := new(openapi.CreateQuoteRequest)
	if err := c.Bind(payload); err != nil {
		return err
	}

	response := openapi.PostQuotes200JSONResponse{
		AuthorId: uuid.New(),
		Text:     "some text here",
	}

	return c.JSON(http.StatusCreated, response)
}
