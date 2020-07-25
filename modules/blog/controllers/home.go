package controllers

import (
	"log"
	"zodream/modules/blog/dao"

	"github.com/kataras/iris/v12"
)

// Index 显示页面
func Index(ctx iris.Context) {
	page, err := ctx.URLParamInt("page")
	if err == nil {
		page = 1
	}
	items, pager, err := dao.GetBlogList(page)
	log.Println("query blog: %d", len(items))
	ctx.ViewData("items", &items)
	ctx.ViewData("pager", &pager)
	ctx.View("blog/index.html")
}
