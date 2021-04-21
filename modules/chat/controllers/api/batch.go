package api

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/chat/dao"
	"zodream.cn/godream/utils/batch"
	"zodream.cn/godream/utils/pagination"
	"zodream.cn/godream/utils/response"
)

func BatchMap(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	var params map[string]interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	userId := c.Keys["user_id"].(uint)
	handleItems := batch.BatchMap{
		"chat/chat": func(data interface{}) interface{} {
			return api.RenderPage(dao.GetHistories(userId), pagination.New(1, 100, 20))
		},
		"chat/user": func(data interface{}) interface{} {
			return dao.GetProfile(userId)
		},
		"chat/friend": func(data interface{}) interface{} {
			return gin.H{
				"data": dao.GetFriendList(userId),
			}
		},
		"chat/group": func(data interface{}) interface{} {
			return gin.H{
				"data": dao.GetGroupList(userId),
			}
		},
	}
	data, err := batch.InvokeBatch(params, handleItems)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.Render(data))
}
