package controllers

import (
	"github.com/kataras/iris"
)

/*
* 显示页面
 */
func Index(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("blog/index.html")
}
