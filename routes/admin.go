package router

import (
	"webProject1/controller"
	"webProject1/utility"
)

// AddAdminRoutes :Adds all admin routes to the app
func AddAdminRoutes(app routeable) {
	app.Use(utility.CheckAdminAccess)
	app.POST("", controller.AddUser)
	app.DELETE("", controller.DeleteUser)
}
