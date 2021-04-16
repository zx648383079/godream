package controllers

import "github.com/gin-gonic/gin"

// Index 显示页面
func Index(ctx *gin.Context) {
	ctx.HTML(200, "chat/index.html", gin.H{})
}
