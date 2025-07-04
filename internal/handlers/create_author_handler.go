package handlers

import (
	"errors"
	"net/http"

	"github.com/GoCodingX/gorilla/internal/repository"
	pkgerrors "github.com/GoCodingX/gorilla/pkg/errors"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/labstack/echo/v4"
)

func (s *QuotesService) PostAuthors(c echo.Context) error {
	// authorization
	if _, err := Authorize(c, PermissionWrite); err != nil {
		return err
	}

	// read the request payload
	createAuthorPayload := new(openapi.CreateAuthorRequest)
	if err := c.Bind(createAuthorPayload); err != nil {
		return err
	}

	// prepare data for persistence
	repoAuthor := toRepoAuthor(createAuthorPayload)

	// persist via repo layer
	err := s.repo.CreateAuthor(c.Request().Context(), repoAuthor)
	if err != nil {
		var errAlreadyExists *repository.AlreadyExistsError
		if errors.As(err, &errAlreadyExists) {
			return pkgerrors.NewEchoErrorResponse(http.StatusConflict, errAlreadyExists.Msg, nil)
		}

		return err
	}

	// prepare http response payload
	response := &openapi.CreateAuthorResponse{
		Id:   repoAuthor.ID,
		Name: repoAuthor.Name,
	}

	// respond
	return c.JSON(http.StatusCreated, response)
}
