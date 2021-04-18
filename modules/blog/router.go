package blog

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/controllers"
	"zodream.cn/godream/modules/blog/controllers/api"
)

func Register(app *gin.RouterGroup) {
	app.GET("/", controllers.Index)
}

func RegisterAPI(app *gin.RouterGroup) {
	app.GET("", api.BlogPage)
	app.GET("tag", api.TagList)
	app.GET("category", api.TermList)
	app.GET("archives", api.ArchivesList)
}
