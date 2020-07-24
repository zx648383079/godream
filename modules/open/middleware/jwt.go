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
			ctx.StopExecution()
			ctx.Values().Get(platform.PlatformKey).(*platform.Platform).RenderFailure(ctx, iris.StatusUnauthorized, "请先登录")
		},

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("zodream cn"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
