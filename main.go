package main

import (
	routes "webProject1/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.AddAdminRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
