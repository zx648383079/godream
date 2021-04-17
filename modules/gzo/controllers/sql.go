package controllers

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/gzo/dao"
	"zodream.cn/godream/utils/response"
)

func SqlTable(ctx *gin.Context) {
	json := ctx.Keys["json"].(response.IJsonResponse)
	ctx.JSON(200, json.RenderData(dao.GetTables()))
}
