package auth

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/auth/controllers"
	"zodream.cn/godream/modules/auth/controllers/api"
)

// Register 注册路由
func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
	app.POST("/login", controllers.Login)
}

// RegisterAPI 注册api路由
func RegisterAPI(app *gin.RouterGroup) {
	app.GET("", api.Index)
	app.GET("/", api.Index)
}
