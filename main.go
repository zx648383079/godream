package main

import (
	"fmt"

	"zodream/configs"
	"zodream/database"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

type app struct {
	*iris.Application
	db *gorm.DB
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
