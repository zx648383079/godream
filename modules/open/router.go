package open

import (
	"zodream/modules/open/controllers"

	"zodream/modules/auth"
	"zodream/modules/chat"
	"zodream/modules/shop"

	"github.com/kataras/iris"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	app.PartyFunc("/auth", auth.Register)
	app.PartyFunc("/chat", chat.Register)
	app.PartyFunc("/shop", shop.Register)
}
