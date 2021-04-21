package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"zodream.cn/godream/configs"
	"zodream.cn/godream/modules/chat/dao"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/utils/response"
)

func MessageList(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var query models.MssageQuery
	c.BindJSON(&query)
	data := dao.GetMessageList(userId, query.StartTime, query.ItemType, query.ItemId)
	c.JSON(200, api.RenderData(data))
}

func Ping(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var query models.MssageQuery
	c.BindJSON(&query)
	data := dao.GetPing(userId, query.StartTime, query.ItemType, query.ItemId)
	c.JSON(200, api.Render(data))
}

func Revoke(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	id, _ := strconv.Atoi(c.Query("id"))
	err := dao.RevokeMessage(userId, uint(id))
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(true))
}

func SendText(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var form models.MssageForm
	c.BindJSON(&form)
	data, err := dao.SendText(userId, form.ItemType, form.ItemId, form.Content)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(data))
}

func SendImage(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var form models.MssageForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	file, url := configs.UploadRandomFileName(f.Filename)
	if err := c.SaveUploadedFile(f, file); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	data, err := dao.SendImage(userId, form.ItemType, form.ItemId, url)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(data))
}

func SendVideo(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var form models.MssageForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	file, url := configs.UploadRandomFileName(f.Filename)
	if err := c.SaveUploadedFile(f, file); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	data, err := dao.SendVideo(userId, form.ItemType, form.ItemId, url)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(data))
}

func SendVoice(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var form models.MssageForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	file, url := configs.UploadRandomFileName(f.Filename)
	if err := c.SaveUploadedFile(f, file); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	data, err := dao.SendVoice(userId, form.ItemType, form.ItemId, url)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(data))
}

func SendFile(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	userId := c.Keys["user_id"].(uint)
	var form models.MssageForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	file, url := configs.UploadRandomFileName(f.Filename)
	if err := c.SaveUploadedFile(f, file); err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	data, err := dao.SendFile(userId, form.ItemType, form.ItemId, f.Filename, url)
	if err != nil {
		c.JSON(400, api.RenderFailure(err))
		return
	}
	c.JSON(200, api.RenderData(data))
}
