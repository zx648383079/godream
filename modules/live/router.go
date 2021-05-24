package live

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/live/controllers/api"
	"zodream.cn/godream/modules/live/rmtp"
	"zodream.cn/godream/modules/open/middleware"
)

func RenderAPI(app *gin.RouterGroup) {
	auth := app.Group("", middleware.JWTAuthorize(true))
	{
		auth.GET("/hls/:name.m3u8", api.M3u8PlayList)
		auth.GET("/hls/:name.ts", api.TsFile)
	}
	setupRmtp(app)
}

func setupRmtp(app *gin.RouterGroup) {
	hub := rmtp.NewServer()
	// go hub.Run()
	app.GET("/rmtp/:channel", func(c *gin.Context) {
		hub.Serve(c.Writer, c.Request, c.Param("channel"))
	})
}
