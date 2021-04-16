package open

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/open/controllers"
)

type (
	GroupFunc    = func(app *gin.RouterGroup)
	GroupFuncMap = map[string]GroupFunc
)

// Register 注册路由
func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
}
