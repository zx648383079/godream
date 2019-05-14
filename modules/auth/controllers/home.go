package controllers

import (
	"github.com/kataras/iris"
)

type LoginForm struct {
	Email      string `form:"email"`
	Password   string
	RememberMe bool
}

func Index(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("auth/index.html")
}

func Login(ctx iris.Context) {
	data := LoginForm{}
	err := ctx.ReadForm(&data)
	ctx.JSON(iris.Map{
		"code":  200,
		"email": data,
		"a":     ctx.FormValue("email"),
		"err":   err,
	})
}
