package handlers

import (
	"errors"
	"net/http"

	"github.com/GoCodingX/gorilla/internal/repository"
	pkgerrors "github.com/GoCodingX/gorilla/pkg/errors"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) PostQuotes(c echo.Context) error {
	// authorization
	user, err := checkPermission(c, PermissionWrite)
	if err != nil {
		return err
	}

	// read the request payload
	createQuotePayload := new(openapi.CreateQuoteRequest)

	err = c.Bind(createQuotePayload)
	if err != nil {
		return err
	}

	// prepare data for persistence
	repoQuote := toRepoQuote(createQuotePayload, user.Username)

	// persist via repo layer
	err = s.repo.CreateQuote(c.Request().Context(), repoQuote)
	if err != nil {
		var errInvalidReferenceError *repository.InvalidReferenceError
		if errors.As(err, &errInvalidReferenceError) {
			return pkgerrors.NewEchoBadRequestResponse(&[]openapi.Detail{{
				Field:   "author_id",
				Message: "no author with such id exists",
			}})
		}

		return err
	}

	// prepare http response payload
	response := openapi.QuoteResponse{
		Id:       repoQuote.ID,
		Text:     repoQuote.Text,
		AuthorId: repoQuote.AuthorID,
	}

	// respond
	return c.JSON(http.StatusCreated, response)
}
