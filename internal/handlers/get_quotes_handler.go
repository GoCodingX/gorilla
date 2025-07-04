package handlers

import (
	"net/http"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) GetQuotes(c echo.Context, params openapi.GetQuotesParams) error {
	// authorization
	_, err := checkPermission(c, PermissionRead)
	if err != nil {
		return err
	}

	// query via the repo layer
	quotes, err := s.repo.GetQuotes(c.Request().Context(), &repository.GetQuotesParams{
		Author: params.Author,
	})
	if err != nil {
		return err
	}

	// prepare http response payload
	responsePayload := make([]openapi.CreateQuoteResponse, len(quotes))
	for i, q := range quotes {
		responsePayload[i] = openapi.CreateQuoteResponse{
			Id:       q.ID,
			Text:     q.Text,
			AuthorId: q.AuthorID,
		}
	}

	// respond
	return c.JSON(http.StatusCreated, responsePayload)
}
