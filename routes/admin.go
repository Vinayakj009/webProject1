package routes

import (
	controller "webProject1/controller"

	echo "github.com/labstack/echo/v4"
)

// AddAdminRoutes :Adds all admin routes to the app
func AddAdminRoutes(app *echo.Echo) {
	app.GET("/", controller.Hello)
}
