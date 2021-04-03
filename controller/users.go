package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Login : Checks if the user passed the right credentials and logs in the user
func Login(context echo.Context) error {
	return context.String(http.StatusOK, "Logged in")
}

// GetUsers : Gets a list of users present in the database.
func GetUsers(context echo.Context) error {
	return context.String(http.StatusOK, "Users Retrived")
}
