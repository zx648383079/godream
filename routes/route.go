package routes

import (
	"github.com/gin-gonic/gin"

	"zodream.cn/godream/configs"
	"zodream.cn/godream/controllers"
	"zodream.cn/godream/modules/auth"
	"zodream.cn/godream/modules/blog"
	"zodream.cn/godream/modules/chat"
	"zodream.cn/godream/modules/gzo"
	"zodream.cn/godream/utils/response"
)

func Register(app *gin.Engine) {
	app.Use(func(ctx *gin.Context) {
		print(ctx.FullPath())
		ctx.Keys["json"] = new(response.JSONResponse)
		ctx.Next()
	})
	app.GET("/", controllers.Index)
	app.GET("/friend_link", controllers.FriendLink)
	app.GET("/about", controllers.About)
	app.GET("/to", controllers.To)
	routes := GroupFuncMap{
		"/auth": auth.Register,
		"/blog": blog.Register,
		"/chat": chat.Register,
	}
	if configs.Config.Debug {
		routes["/gzo"] = gzo.Register
	}
	for path, v := range routes {
		g := app.Group(path)
		v(g)
	}
}
