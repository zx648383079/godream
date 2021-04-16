package controllers

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/utils"
)

// Index 显示页面
func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

// About 关于我们
func About(ctx *gin.Context) {
	ctx.HTML(200, "home/about.html", gin.H{})
}

// FriendLink 友情链接
func FriendLink(ctx *gin.Context) {
	ctx.HTML(200, "home/friend_link.html", gin.H{})
}

// To 跳转到外部链接
func To(ctx *gin.Context) {
	uri := ctx.GetString("url")
	if uri != "" {
		uri = utils.Base64Decode(uri + "=")
	}
	if uri == "" {
		uri = "/"
	}
	ctx.HTML(200, "home/to.html", gin.H{
		"url": uri,
	})
}
