package main

import (
	router "webProject1/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Create groups
	adminGroup := app.Group("/user")
	usersGroup := app.Group("/users")

	// Routes
	router.AddAdminRoutes(adminGroup)
	router.AddUsersRoutes(usersGroup)

	// Start server
	app.Logger.Fatal(app.Start(":1323"))
}
