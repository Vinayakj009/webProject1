package main

import (
	routes "webProject1/routes"

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
	adminGroup := app.Group("/admin")

	// Routes
	routes.AddAdminRoutes(adminGroup)

	// Start server
	app.Logger.Fatal(app.Start(":1323"))
}
