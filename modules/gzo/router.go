package gzo

import (
	"zodream/modules/gzo/controllers"

	"github.com/kataras/iris"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
}
