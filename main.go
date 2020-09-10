package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"

	"os"
	"zodream/configs"
	"zodream/controllers"
	"zodream/database"
	"zodream/modules/auth"
	"zodream/modules/blog"
	"zodream/modules/chat"
	"zodream/modules/gzo"
	"zodream/modules/open"
	"zodream/modules/shop"
	"zodream/utils/response"
)

type app struct {
	*iris.Application
	db *gorm.DB
}

func (app *app) Register() {
	app.Use(func(ctx context.Context) {
		print(ctx.Path())
		ctx.Values().Set("json", new(response.JSONResponse))
		ctx.Next()
	})
	homeRoute := app.Get("/", controllers.Index)
	homeRoute.Name = "home"
	app.Get("/home", controllers.Index)
	aboutRoute := app.Get("/about", controllers.About)
	aboutRoute.Name = "abount"
	linkRoute := app.Get("/friend_link", controllers.FriendLink)
	linkRoute.Name = "friend_link"
	app.PartyFunc("/auth", auth.Register)
	app.PartyFunc("/blog", blog.Register)
	app.PartyFunc("/chat", chat.Register)
	app.PartyFunc("/shop", shop.Register)
	app.PartyFunc("/open", open.Register)
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
	rv := router.NewRoutePathReverser(app, router.WithHost(configs.Host()))
	tmpl.AddFunc("url", rv.URL)
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl.Layout("layouts/layout.html"))
	fmt.Println(configs.Config.Favicon)
	app.Favicon(configs.Config.Favicon)
	app.HandleDir("/assets", configs.Config.Asset)

	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.Writef("404 not found here")
	})
	app.Run(iris.Addr(":" + configs.Config.Port))
}
