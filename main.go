package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"zodream.cn/godream/configs"
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/open/middleware"
	"zodream.cn/godream/routes"
	"zodream.cn/godream/sessions"
	"zodream.cn/godream/view"
)

func main() {
	configs.Init("app.json")
	if configs.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	database.InitDb()
	r := gin.Default()
	r.Use(middleware.CORS)
	r.Use(sessions.New())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routes.Register(r)
	routes.RegisterAPI(r)
	r.Static("/assets", configs.Config.Asset)
	r.StaticFile("/favicon.ico", configs.Config.Favicon)
	r.SetFuncMap(view.GenerateFuns)
	r.LoadHTMLGlob(configs.Config.View)
	r.Run(":" + configs.Config.Port)
}
