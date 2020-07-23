package open

import (
	"zodream/modules/open/controllers"

	"zodream/modules/auth"

	"github.com/kataras/iris"
)

// Register 注册路由
func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	app.PartyFunc("/auth", auth.RegisterAPI)
	// app.PartyFunc("/chat", chat.RegisterAPI)
	// app.PartyFunc("/shop", shop.RegisterAPI)
}
