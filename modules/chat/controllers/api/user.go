package api

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/chat/dao"
)

func UserProfile(c *gin.Context) {
	// api := c.Keys["json"].(response.IJsonResponse)
	user := dao.GetProfile(c.Keys["user_id"].(uint))
	c.JSON(200, user)
}
