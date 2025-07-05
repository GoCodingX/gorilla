package handlers

import (
	"net/http"
	"time"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) GetQuotes(c echo.Context, params openapi.GetQuotesParams) error {
	// authorization
	_, err := Authorize(c, PermissionRead)
	if err != nil {
		return err
	}

	// query via the repo layer
	repoQuotes, cursor, err := s.repo.GetQuotes(c.Request().Context(), &repository.GetQuotesParams{
		Author:          params.Author,
		CursorCreatedAt: params.CursorCreatedAt,
		CursorID:        params.CursorId,
	})
	if err != nil {
		return err
	}

	// prepare http response payload
	quotes := make([]openapi.QuoteResponse, len(repoQuotes))
	for i, q := range repoQuotes {
		quotes[i] = openapi.QuoteResponse{
			Id:       q.ID,
			Text:     q.Text,
			AuthorId: q.AuthorID,
		}
	}

	var nextCursor *openapi.NextCursor
	if cursor != nil {
		nextCursor = &openapi.NextCursor{
			CreatedAt: cursor.CreatedAt.Format(time.RFC3339),
			Id:        cursor.ID.String(),
		}
	}

	// prepare http response payload
	response := &openapi.GetQuotesResponse{
		Quotes:     quotes,
		NextCursor: nextCursor,
	}

	// respond
	return c.JSON(http.StatusCreated, response)
}
