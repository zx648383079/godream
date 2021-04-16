package chat

import (
	"zodream.cn/godream/modules/chat/controllers"
	"zodream.cn/godream/modules/chat/server"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
	setupWebsocket(app)
}

func setupWebsocket(app *gin.RouterGroup) {
	hub := server.NewHub()
	go hub.Run()
	app.GET("/ws", func(c *gin.Context) {
		server.ServeWs(hub, c)
	})
}
