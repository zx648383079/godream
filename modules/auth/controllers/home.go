package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/auth/models"
)

// LoginForm 登录表单
type LoginForm struct {
	Email      string `form:"email"`
	Password   string
	RememberMe bool
}

// Index 登录页面
func Index(ctx *gin.Context) {
	ctx.HTML(200, "auth/index.html", gin.H{})
}

// Login 登录
func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user, err := models.LoginEmail(ctx.PostForm("email"), ctx.PostForm("password"))
	if err == nil {
		session.Set("userID", user.ID)
		session.Save()
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": user,
		"err":  err,
	})
}
