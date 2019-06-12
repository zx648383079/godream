package controllers

import (
	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"messageL": "a",
		"b":        "1",
	})
}
