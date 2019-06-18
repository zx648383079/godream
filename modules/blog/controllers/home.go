package controllers

import (
	"github.com/kataras/iris"
	"zodream/modules/blog/models"
	"log"
)

// Index 显示页面
func Index(ctx iris.Context) {
	blogs := models.GetBlogList()
	log.Println("query blog: %d", len(blogs))
	ctx.ViewData("Blogs", &blogs)
	ctx.View("blog/index.html")
}
