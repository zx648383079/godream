package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/dao"
	"zodream.cn/godream/modules/blog/models"
	"zodream.cn/godream/utils/response"
)

func BlogPage(c *gin.Context) {
	_, exist := c.GetQuery("id")
	if exist {
		BlogDetail(c)
		return
	}
	var queries models.BlogQueries
	c.ShouldBindQuery(queries)
	items, pager, _ := dao.GetBlogList(&queries)
	api := c.Keys["json"].(response.IJsonResponse)
	c.JSON(200, api.RenderPage(items, pager))
}

func BlogDetail(c *gin.Context) {
	val, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(val)
	data, err := dao.GetBlogFull(id)
	api := c.Keys["json"].(response.IJsonResponse)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(400, api.Render(data))
}
