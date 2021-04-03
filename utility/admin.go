package utility

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CheckAdminAccess is a middleware function which check whether the authorisation request header has the right value.
// If the request has the correct authorisation header, then the request is allowed. Otherwise the request is not allowed
func CheckAdminAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header
		if header.Get("random_string") == "tasty" {
			return next(context)
		}
		return context.String(http.StatusForbidden, "Your are not an authorised admin")
	}
}
