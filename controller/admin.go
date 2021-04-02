package controller

import (
	"net/http"

	echo "github.com/labstack/echo"
)

// Hello : Prints hello, world to the api
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "This is the admin base")
}
