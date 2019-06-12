package main

import (
	"fmt"

	"zodream/configs"
	"zodream/database"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"

	"os"
	"zodream/controllers"
	"zodream/modules/auth"
	"zodream/modules/chat"
	"zodream/modules/gzo"
	"zodream/modules/shop"
)

type app struct {
	*iris.Application
	db *gorm.DB
}

func (app *app) Register() {
	app.Get("/", controllers.Index)
	app.Get("/home", controllers.Index)
	app.PartyFunc("/auth", auth.Register)
	app.PartyFunc("/chat", chat.Register)
	app.PartyFunc("/shop", shop.Register)
	if os.Getenv("DEBUG") == "true" {
		app.PartyFunc("/gzo", gzo.Register)
	}
}

func main() {
	configs.Init("app.json")
	app := &app{iris.Default(), database.New()}
	defer app.db.Close()
	app.Register()
	tmpl := iris.HTML(configs.Config.View, ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl.Layout("layout.html"))
	fmt.Println(configs.Config.Favicon)
	app.Favicon(configs.Config.Favicon)
	app.StaticWeb("/assets", configs.Config.Asset)

	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.Writef("404 not found here")
	})
	app.Run(iris.Addr(":" + configs.Config.Port))
}
