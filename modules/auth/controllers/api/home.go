package api

import (
	"zodream/utils/response"

	"github.com/kataras/iris/v12"
)

// Index 登录页面
func Index(ctx iris.Context) {

	api := ctx.Values().Get("json").(response.IJsonResponse)
	ctx.JSON(api.RenderData(1))
}
