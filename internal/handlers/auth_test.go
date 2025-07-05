package handlers_test

import (
	"net/http"
	"testing"

	"github.com/GoCodingX/gorilla/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	t.Run("returns echo.ErrUnauthorized when user is not set in context", func(t *testing.T) {
		// create an echo ctx
		c, _ := newEchoContext(nil)

		// act
		_, err := handlers.Authorize(c, handlers.PermissionRead)

		// assert
		assert.Error(t, err)
		var httpErr *echo.HTTPError
		if assert.ErrorAs(t, err, &httpErr) {
			assert.Equal(t, http.StatusUnauthorized, httpErr.Code)
			assert.Equal(t, "Unauthorized", httpErr.Message)
		}
	})

	t.Run("returns echo.ErrUnauthorized when user is not set as the right type in context", func(t *testing.T) {
		// create an echo ctx
		c, _ := newEchoContext(nil)

		// set user to a different type
		c.Set("user", "user_value")

		// act
		_, err := handlers.Authorize(c, handlers.PermissionRead)

		// assert
		assert.Error(t, err)
		var httpErr *echo.HTTPError
		if assert.ErrorAs(t, err, &httpErr) {
			assert.Equal(t, http.StatusUnauthorized, httpErr.Code)
			assert.Equal(t, "Unauthorized", httpErr.Message)
		}
	})

	t.Run("returns echo.ErrForbidden when user has 'read' permission and 'write' is required", func(t *testing.T) {
		// create an echo ctx
		c, _ := newEchoContext(nil)

		// set user
		c.Set("user", &handlers.User{
			Permission: handlers.PermissionRead,
		})

		// act
		_, err := handlers.Authorize(c, handlers.PermissionWrite)

		// assert
		assert.Error(t, err)
		var httpErr *echo.HTTPError
		if assert.ErrorAs(t, err, &httpErr) {
			assert.Equal(t, http.StatusForbidden, httpErr.Code)
			assert.Equal(t, "Forbidden", httpErr.Message)
		}
	})

	t.Run("returns no error when user has the right permission", func(t *testing.T) {
		// create an echo ctx
		c, _ := newEchoContext(nil)

		// set user
		username := "some_username"
		password := "some_password"
		c.Set("user", &handlers.User{
			Username:   username,
			Password:   password,
			Permission: handlers.PermissionWrite,
		})

		// act
		user, err := handlers.Authorize(c, handlers.PermissionRead)

		// assert
		assert.NoError(t, err)
		if assert.NotNil(t, user) {
			assert.Equal(t, username, user.Username)
			assert.Equal(t, handlers.PermissionWrite, user.Permission)
			assert.Equal(t, password, user.Password)
		}
	})
}
