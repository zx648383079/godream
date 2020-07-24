package open

import (
	"zodream/modules/auth"
	"zodream/modules/open/controllers"
	"zodream/modules/open/middleware"

	"github.com/kataras/iris/v12"
)

// Register 注册路由
func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	app.Use(middleware.CORS, middleware.RESTful)
	{
		app.PartyFunc("/auth", auth.RegisterAPI)
		// app.PartyFunc("/chat", chat.RegisterAPI)
		// app.PartyFunc("/shop", shop.RegisterAPI)
	}

}
