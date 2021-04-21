package chat

import (
	"zodream.cn/godream/modules/chat/controllers"
	"zodream.cn/godream/modules/chat/controllers/api"
	"zodream.cn/godream/modules/chat/server"
	"zodream.cn/godream/modules/open/middleware"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
	setupWebsocket(app)
}

func RegisterAPI(app *gin.RouterGroup) {
	auth := app.Group("", middleware.JWTAuthorize(true))
	{
		auth.POST("/batch", api.BatchMap)
		auth.POST("/message", api.MessageList)
		auth.POST("/message/ping", api.Ping)
		auth.POST("/message/send_text", api.SendText)
		auth.POST("/message/send_image", api.SendImage)
		auth.POST("/message/send_video", api.SendVideo)
		auth.POST("/message/send_voice", api.SendVoice)
		auth.POST("/message/send_file", api.SendFile)
		auth.DELETE("/message/revoke", api.Revoke)
		auth.DELETE("/chat/remove", api.HistoryRemove)
		auth.GET("/chat", api.Histories)
	}
	setupWebsocket(app)
}

func setupWebsocket(app *gin.RouterGroup) {
	hub := server.NewHub()
	go hub.Run()
	app.GET("/ws", func(c *gin.Context) {
		server.ServeWs(hub, c)
	})
}
