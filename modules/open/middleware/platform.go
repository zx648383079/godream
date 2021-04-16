package middleware

import (
	"zodream.cn/godream/modules/open/platform"

	"github.com/gin-gonic/gin"
)

// REST RESTful
var REST gin.HandlerFunc

// RESTful rest 代码
type RESTful struct {
}

// New 初始化
func New() *RESTful {
	rest := new(RESTful)
	return rest
}

// Serve 执行
func (rest *RESTful) Serve(ctx *gin.Context) {
	appid := ctx.GetString("appid")
	json := new(platform.PlatformResponse)
	api, err := platform.NewPlatform(appid)
	if err != nil {
		ctx.AbortWithStatusJSON(400, json.RenderFailure(err.Error()))
		return
	}
	if !api.VerifyRule(ctx.FullPath()) {
		ctx.AbortWithStatusJSON(400, json.RenderFailure(err.Error()))
		return
	}
	err = api.Verify(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, json.RenderFailure(err.Error()))
		return
	}
	json.Platform = api
	ctx.Keys["json"] = json
	ctx.Next()
}

func initREST() {
	REST = New().Serve
}
