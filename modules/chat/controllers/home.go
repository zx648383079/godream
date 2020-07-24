package controllers

import (
	"github.com/kataras/iris/v12"
)

// Index 显示页面
func Index(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("chat/index.html")
}
