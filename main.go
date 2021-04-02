package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vinayakj009/webProject1/routes"
)

func main() {
	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Create groups
	adminGroup := app.Group("/admin")
	other := app.Group("/other")

	// Routes
	routes.AddAdminRoutes(adminGroup)
	routes.AddAdminRoutes(other)
	routes.AddAdminRoutes(app)

	// Start server
	app.Logger.Fatal(app.Start(":1323"))
}
