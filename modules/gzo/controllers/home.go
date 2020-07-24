package controllers

import (
	"github.com/kataras/iris/v12"
)

func Index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"hh": "hh",
	})
}
