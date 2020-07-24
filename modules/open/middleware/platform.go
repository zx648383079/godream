package middleware

import (
	"zodream/modules/open/platform"
	"zodream/utils"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/context"
)

// RESTful rest 代码
type RESTful struct {
}

// New 初始化
func New() *RESTful {
	rest := new(RESTful)
	return rest
}

// Serve 执行
func (rest *RESTful) Serve(ctx context.Context) {
	appid := ctx.URLParam("appid")
	api, err := platform.NewPlatform(appid)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.FailureJson(err.Error))
		ctx.StopExecution()
		return
	}
	err = api.Verify(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.FailureJson(err.Error))
		ctx.StopExecution()
		return
	}
	ctx.Values().Set(platform.PlatformKey, api)
	ctx.Next()
}
