package middleware

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/configs"
	"zodream.cn/godream/modules/open/platform"
)

var (
	// JWT JWT Middleware
	JWT gin.HandlerFunc
)

func initJWT() {
	JWT = func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(configs.Config.Auth.Key)
			return b, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, c.Keys["json"].(*platform.PlatformResponse).RenderFailure("请先登录"))
			return
		}
		c.Next()
	}

}
