package middleware

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/configs"
	"zodream.cn/godream/utils/response"
)

func JWTAuthorize(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		api := c.Keys["json"].(response.IJsonResponse)
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(configs.Config.Auth.Key)
			return b, nil
		})
		if err != nil {
			if required {
				c.AbortWithStatusJSON(401, api.RenderFailure("请先登录"))
			}
			return
		}

		claims := token.Claims.(jwt_lib.MapClaims)
		if err := claims.Valid(); err != nil {
			if required {
				c.AbortWithStatusJSON(401, api.RenderFailure("请先登录"))
			}
			return
		}
		c.Keys["user_id"] = uint(claims["sub"].(float64))
		c.Next()
	}
}
