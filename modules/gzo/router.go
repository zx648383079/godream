package gzo

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/gzo/controllers"
)

func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
	app.GET("/sql/table", controllers.SqlTable)
	app.POST("/template/model", controllers.TplModel)
}
