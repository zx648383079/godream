package api

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/dao"
	"zodream.cn/godream/modules/blog/models"
	"zodream.cn/godream/utils/response"
)

func BlogPage(c *gin.Context) {
	var queries models.BlogQueries
	c.ShouldBindQuery(queries)
	items, pager, _ := dao.GetBlogList(&queries)
	api := c.Keys["json"].(response.IJsonResponse)
	c.JSON(200, api.RenderPage(items, pager))
}
