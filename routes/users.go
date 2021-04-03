package router

import (
	"webProject1/controller"
	"webProject1/utility"
)

// AddUsersRoutes :Adds all admin routes to the app
func AddUsersRoutes(app routeable) {
	app.POST("/login", controller.Login)
	app.GET("", controller.GetUsers, utility.CheckLoggedIn)
}
