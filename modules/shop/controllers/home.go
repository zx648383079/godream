package controllers

import (
	"github.com/kataras/iris/v12"
)

func Index(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("shop/index.html")
}
