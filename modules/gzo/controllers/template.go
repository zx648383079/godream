package controllers

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/gzo/dao"
	"zodream.cn/godream/utils/response"
)

func TplModel(ctx *gin.Context) {
	table := ctx.PostForm("table")
	json := ctx.Keys["json"].(response.IJsonResponse)
	ctx.JSON(200, json.RenderData(gin.H{
		"code": dao.RenderModel(table),
	}))
}
