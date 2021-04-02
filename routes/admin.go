package routes

import (
	"github.com/vinayakj009/webProject1/controller"

	"github.com/labstack/echo"
)

// AddAdminRoutes :Adds all admin routes to the app
func AddAdminRoutes(app echo.Routeable) {
	app.GET("", controller.Hello)
}
