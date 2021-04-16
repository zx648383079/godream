package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/configs"
	"zodream.cn/godream/utils/response"
)

// Index 登录页面
func Index(ctx *gin.Context) {

	api := ctx.Keys["json"].(response.IJsonResponse)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	ss, err := token.SignedString([]byte(configs.Config.Auth.Key))
	if err != nil {
		ctx.JSON(400, api.RenderFailure(err))
		return
	}
	ctx.JSON(200, api.RenderData(ss))
}
