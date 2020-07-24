package wx

import (
	"zodream/modules/wx/controllers"

	"github.com/kataras/iris/v12"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
}
