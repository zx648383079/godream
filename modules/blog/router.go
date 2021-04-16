package blog

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/controllers"
)

func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
}
