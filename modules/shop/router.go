package shop

import (
	"zodream/modules/shop/controllers"

	"github.com/kataras/iris"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	app.Post("/login", controllers.Login)
}
