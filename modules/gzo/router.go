package gzo

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/gzo/controllers"
)

func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
}
