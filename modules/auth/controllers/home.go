package controllers

import (
	"zodream/modules/auth/models"

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
	user, err := models.LoginEmail(ctx.FormValue("email"), ctx.FormValue("password"))
	ctx.JSON(iris.Map{
		"code": 200,
		"data": user,
		"err":  err,
	})
}
