package api

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/blog/dao"
	"zodream.cn/godream/utils/response"
)

func TagList(c *gin.Context) {
	items := dao.GetTags()
	api := c.Keys["json"].(response.IJsonResponse)
	c.JSON(200, api.RenderData(items))
}
