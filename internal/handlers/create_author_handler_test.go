package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/gen/openapi"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPostAuthors(t *testing.T) {
	t.Run("returns Forbidden when user does not have the right permission", func(t *testing.T) {
		// create echo context
		c, _ := newEchoContext(&newEchoContextParams{
			method:  http.MethodPost,
			target:  "/authors",
			payload: `{}`,
		})

		// set a user without the right permissions in context
		c.Set("user", &handlers.User{
			Permission: handlers.PermissionRead,
		})

		// create service
		svc, _ := newServiceWithMockRepo(t)

		// act
		err := svc.PostAuthors(c)

		// assert
		if assert.Error(t, err) {
			var httpErr *echo.HTTPError
			if assert.ErrorAs(t, err, &httpErr) {
				assert.Equal(t, http.StatusForbidden, httpErr.Code)
				assert.Equal(t, "Forbidden", httpErr.Message)
			}
		}
	})

	t.Run("returns StatusConflict when author already exists", func(t *testing.T) {
		reqPayload := `{"name":"Dan Brown"}`

		// create echo context
		c, _ := newEchoContext(&newEchoContextParams{
			method:  http.MethodPost,
			target:  "/authors",
			payload: reqPayload,
		})

		// set a user without the right permissions in context
		c.Set("user", &handlers.User{
			Permission: handlers.PermissionWrite,
		})

		// create service
		svc, repo := newServiceWithMockRepo(t)

		repo.
			EXPECT().
			CreateAuthor(c.Request().Context(), gomock.Any()).
			Return(&repository.AlreadyExistsError{})

		// act
		err := svc.PostAuthors(c)

		// assert
		if assert.Error(t, err) {
			var httpErr *echo.HTTPError
			if assert.ErrorAs(t, err, &httpErr) {
				errRsp, ok := httpErr.Message.(*openapi.ErrorResponse)
				if assert.True(t, ok) {
					assert.Equal(t, http.StatusConflict, errRsp.Code)
					assert.Equal(t, "Conflict", errRsp.Status)
				}
			}
		}
	})

	t.Run("returns StatusCreated", func(t *testing.T) {
		authorName := "Dan Brown"
		reqPayload := fmt.Sprintf(`{"name":"%s"}`, authorName)

		// create echo context
		c, rec := newEchoContext(&newEchoContextParams{
			method:  http.MethodPost,
			target:  "/authors",
			payload: reqPayload,
		})

		// set a user without the right permissions in context
		c.Set("user", &handlers.User{
			Permission: handlers.PermissionWrite,
		})

		// create service
		svc, repo := newServiceWithMockRepo(t)

		repo.
			EXPECT().
			CreateAuthor(c.Request().Context(), gomock.Any()).
			Return(nil)

		// act
		err := svc.PostAuthors(c)

		// assert
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var resp *openapi.CreateAuthorResponse
			err = json.Unmarshal(rec.Body.Bytes(), &resp)
			if assert.NoError(t, err) {
				assert.Equal(t, authorName, resp.Name)
				assert.NotEqual(t, uuid.Nil, resp.Id)
			}
		}
	})
}
