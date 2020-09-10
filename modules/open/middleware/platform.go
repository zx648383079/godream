package middleware

import (
	"zodream/modules/open/platform"

	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/context"
)

// REST RESTful
var REST context.Handler

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
	json := new(platform.PlatformResponse)
	api, err := platform.NewPlatform(appid)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(json.RenderFailure(err.Error()))
		ctx.StopExecution()
		return
	}
	if !api.VerifyRule(ctx.Path()) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(json.RenderFailure(err.Error()))
		ctx.StopExecution()
		return
	}
	err = api.Verify(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(json.RenderFailure(err.Error()))
		ctx.StopExecution()
		return
	}
	json.Platform = api
	ctx.Values().Set("json", json)
	ctx.Next()
}

func initREST() {
	REST = New().Serve
}
