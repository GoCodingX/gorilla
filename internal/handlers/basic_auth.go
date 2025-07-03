package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var basicAuthMiddleware = middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	if u, ok := users[username]; ok && u.Password == password {
		method := c.Request().Method
		reqUri := c.Request().RequestURI
		urlPath := c.Request().URL.Path
		fmt.Println(method, reqUri, urlPath)

		// password is no longer needed after this point, get rid of it for higher security
		// shallow copy so that the password is not wiped for future calls
		userCopy := *u
		userCopy.Password = ""

		c.Set("user", &userCopy)

		return true, nil
	}

	return false, nil
})
