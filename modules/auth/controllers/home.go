package controllers

import (
	"zodream/modules/auth/models"
	"zodream/sessions"

	"github.com/kataras/iris/v12"
)

// LoginForm 登录表单
type LoginForm struct {
	Email      string `form:"email"`
	Password   string
	RememberMe bool
}

// Index 登录页面
func Index(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("auth/index.html")
}

// Login 登录
func Login(ctx iris.Context) {
	session := sessions.Driver.Start(ctx)
	user, err := models.LoginEmail(ctx.FormValue("email"), ctx.FormValue("password"))
	if err == nil {
		session.Set("userID", user.ID)
	}
	ctx.JSON(iris.Map{
		"code": 200,
		"data": user,
		"err":  err,
	})
}
