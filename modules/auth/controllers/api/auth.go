package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/configs"
	"zodream.cn/godream/modules/auth/dao"
	"zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/utils/response"
)

// Index 登录页面
func Login(ctx *gin.Context) {
	api := ctx.Keys["json"].(response.IJsonResponse)
	var form models.LoginEmail
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(400, api.RenderFailure(err))
		return
	}
	user, err := dao.LoginEmail(form)
	if err != nil {
		ctx.JSON(400, api.RenderFailure(err))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	ss, err := token.SignedString([]byte(configs.Config.Auth.Key))
	if err != nil {
		ctx.JSON(400, api.RenderFailure(err))
		return
	}
	ctx.JSON(200, api.Render(models.ParseToken(user, ss)))
}
