package blog

import (
	"zodream/modules/blog/controllers"

	"github.com/kataras/iris"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
}
