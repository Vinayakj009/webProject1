package admin

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Hello : Prints hello, world to the api
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
