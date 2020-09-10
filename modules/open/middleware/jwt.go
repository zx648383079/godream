package middleware

import (
	"zodream/modules/open/platform"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var (
	// JWT JWT Middleware
	JWT *jwt.Middleware
)

func initJWT() {
	JWT = jwt.New(jwt.Config{
		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			}
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(ctx.Values().Get("json").(*platform.PlatformResponse).RenderFailure("请先登录"))
			ctx.StopExecution()
		},

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("zodream cn"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
