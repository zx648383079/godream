package api

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/contact/dao"
	"zodream.cn/godream/modules/contact/models"
	"zodream.cn/godream/utils/response"
)

func FriendLink(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	c.JSON(200, api.RenderData(dao.GetLinkList()))
}

func FriendLinkApply(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	var link models.LinkForm
	if err := c.ShouldBind(&link); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	if err := dao.ApplyLink(&link, c.GetInt("user_id")); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(true))
}

func FeedbackSave(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	var form models.FeedbackForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	if err := dao.SaveFeedback(&form, c.GetInt("user_id")); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(true))
}

func Subscribe(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	if err := dao.Subscribe(c.PostForm("email")); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(true))
}

func Unsubscribe(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	if err := dao.Unsubscribe(c.PostForm("email")); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(true))
}
