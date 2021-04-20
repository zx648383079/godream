package contact

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/contact/controllers/api"
	"zodream.cn/godream/modules/open/middleware"
)

func RegisterAPI(app *gin.RouterGroup) {
	app.Use(middleware.JWTAuthorize(false))
	app.GET("friend_link", api.FriendLink)
	app.POST("friend_link/apply", api.FriendLinkApply)
	app.POST("feedback", api.FeedbackSave)
	app.POST("subscribe", api.Subscribe)
	app.POST("unsubscribe", api.Unsubscribe)
}
