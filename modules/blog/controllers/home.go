package controllers

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/dao"
	"zodream.cn/godream/modules/blog/models"
)

// Index 显示页面
func Index(ctx *gin.Context) {
	var queries models.BlogQueries
	ctx.ShouldBindQuery(queries)
	items, pager, _ := dao.GetBlogList(&queries)
	ctx.HTML(200, "blog/index.html", gin.H{
		"items": &items,
		"pager": &pager,
	})
}
