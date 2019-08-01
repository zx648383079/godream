package controllers

import (
	"github.com/kataras/iris"
)

// Index 显示页面
func Index(ctx iris.Context) {
	ctx.View("home/index.html")
}

// About 关于我们
func About(ctx iris.Context) {
	ctx.View("home/about.html")
}

// FriendLink 友情链接
func FriendLink(ctx iris.Context) {
	ctx.View("home/friend_link.html")
}

// To 跳转到外部链接
func To(ctx iris.Context) {
	uri := ctx.URLParam("url")
	ctx.View("home/to.html")
}