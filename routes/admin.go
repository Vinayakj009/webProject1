package admin

import (
	controller "webProject1/controller"

	echo "github.com/labstack/echo/v4"
)

type routeable interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// AddAdminRoutes :Adds all admin routes to the app
func AddAdminRoutes(app routeable) {
	app.GET("", controller.Hello)
}
