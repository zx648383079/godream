package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/dao"
)

// Index 显示页面
func Index(ctx *gin.Context) {
	page := ctx.GetInt("page")
	if page < 1 {
		page = 1
	}
	items, pager, _ := dao.GetBlogList(page)
	log.Println("query blog: %d", len(items))
	ctx.HTML(200, "blog/index.html", gin.H{
		"items": &items,
		"pager": &pager,
	})
}
