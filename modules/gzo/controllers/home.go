package controllers

import (
	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"hh": "hh",
	})
}
