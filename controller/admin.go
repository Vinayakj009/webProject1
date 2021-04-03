package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// AddUser : Adds a user to the database
func AddUser(context echo.Context) error {
	return context.String(http.StatusOK, "User added")
}

// DeleteUser : Deletes a user from the database
func DeleteUser(context echo.Context) error {
	return context.String(http.StatusOK, "User deleted")
}
