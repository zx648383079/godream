package api

import (
	"zodream/modules/open/platform"

	"github.com/kataras/iris/v12"
)

// Index 登录页面
func Index(ctx iris.Context) {

	api := ctx.Values().Get(platform.PlatformKey).(*platform.Platform)
	api.RenderData(ctx, 1)
}
