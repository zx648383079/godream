package router

import (
	"os"
	"zodream/controllers"
	"zodream/modules/auth"
	"zodream/modules/gzo"

	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	app.Get("/", controllers.Index)
	app.Get("/home", controllers.Index)
	app.PartyFunc("/auth", auth.Register)
	if os.Getenv("DEBUG") == "true" {
		app.PartyFunc("/gzo", gzo.Register)
	}
}
