package utility

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CheckLoggedIn is a middleware function which check whether the request has been made by a valid logged in user
func CheckLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header
		if header.Get(echo.HeaderAuthorization) == "Auth_Test" {
			return next(context)
		}
		return context.String(http.StatusForbidden, "You are not logged in")
	}
}
