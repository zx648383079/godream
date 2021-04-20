package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/auth/dao"
	"zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/utils/response"
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
	json := ctx.Keys["json"].(response.IJsonResponse)
	var form models.LoginEmail
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(200, json.RenderFailure(err))
	}
	user, err := dao.LoginEmail(form)
	if err == nil {
		session.Set("user_id", user.ID)
		session.Save()
	}
	ctx.JSON(200, json.RenderData(user))
}
