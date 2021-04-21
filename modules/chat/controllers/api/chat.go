package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/chat/dao"
	"zodream.cn/godream/utils/response"
)

func Histories(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	data := dao.GetHistories(userId)
	c.JSON(200, api.RenderData(data))
}

func HistoryRemove(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	id, _ := strconv.Atoi(c.Query("id"))
	dao.RemoveIdHistory(userId, uint(id))
	c.JSON(200, api.RenderData(true))
}

func FriendList(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	data := dao.GetFriendList(userId)
	c.JSON(200, api.RenderData(data))
}

func GroupList(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	data := dao.GetGroupList(userId)
	c.JSON(200, api.RenderData(data))
}

func FollowUser(c *gin.Context) {

}

func RemoveFriend(c *gin.Context) {

}

func MoveFriend(c *gin.Context) {

}

func SearchUser(c *gin.Context) {

}

func SearchGroup(c *gin.Context) {

}

func ApplyGroup(c *gin.Context) {

}

func CreateGroup(c *gin.Context) {

}

func DisbandGroup(c *gin.Context) {

}

func GroupDetail(c *gin.Context) {

}

func ApplyList(c *gin.Context) {

}

func ApplyAction(c *gin.Context) {

}
