package auth

import (
	"zodream/modules/auth/controllers"

	"github.com/kataras/iris"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
}
