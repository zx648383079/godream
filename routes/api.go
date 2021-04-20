package routes

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/auth"
	"zodream.cn/godream/modules/blog"
	"zodream.cn/godream/modules/chat"
	"zodream.cn/godream/modules/contact"
	"zodream.cn/godream/modules/open"
	"zodream.cn/godream/modules/open/middleware"
)

func RegisterAPI(app *gin.Engine) {
	g := app.Group("/open")
	open.Register(g)
	api := g.Group("")
	api.Use(middleware.REST, middleware.CORS)
	{
		routes := GroupFuncMap{
			"/auth":    auth.RegisterAPI,
			"/blog":    blog.RegisterAPI,
			"/contact": contact.RegisterAPI,
			"/chat":    chat.RegisterAPI,
		}
		for path, v := range routes {
			gr := api.Group(path)
			v(gr)
		}
	}

}
