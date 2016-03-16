package routers

import (
	"github.com/astaxie/beego"
	"zodream/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
