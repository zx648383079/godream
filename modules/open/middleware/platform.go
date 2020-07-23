package middleware

import (
	"zodream/modules/open/models"
	"zodream/utils"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
)

type Middleware struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Middleware {
	jwt := new(Middleware)
	jwt.db = db
	return jwt
}

func (jwt *Middleware) Serve(ctx context.Context) {
	appid := ctx.URLParam("appid")
	platform := models.FindPlatform(appid)
	if platform == nil {
		ctx.StatusCode(404)
		ctx.JSON(utils.FailureJson("APP ID error"))
		ctx.StopExecution()
		return
	}
	ctx.Next()
}
