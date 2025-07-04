package handlers

import (
	"errors"
	"log/slog"

	"github.com/GoCodingX/gorilla/pkg/logger"
	"github.com/labstack/echo/v4"
)

func Authorize(c echo.Context, requiredPermission Permission) (*User, error) {
	user, err := getUserFromContext(c)
	if err != nil {
		logger.Error("failed to get user from context", slog.String("err", err.Error()))

		return nil, echo.ErrUnauthorized
	}

	if user.Permission == PermissionWrite || (user.Permission == PermissionRead && requiredPermission == PermissionRead) {
		return user, nil
	}

	return nil, echo.ErrForbidden
}

func getUserFromContext(c echo.Context) (*User, error) {
	userInterface := c.Get("user")
	if userInterface == nil {
		return nil, errors.New("failed to get property user from context")
	}

	user, ok := userInterface.(*User)
	if !ok {
		return nil, errors.New("failed to assert user to desired type")
	}

	return user, nil
}
