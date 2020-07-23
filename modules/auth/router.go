package auth

import (
	"zodream/modules/auth/controllers"

	"github.com/kataras/iris"
)

// Register 注册路由
func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	app.Post("/login", controllers.Login)
}

// RegisterAPI 注册api路由
func RegisterAPI(app iris.Party) {
	app.Get("/", controllers.Index)
	app.Post("/login", controllers.Login)
}
