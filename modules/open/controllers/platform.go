package controllers

import (
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.HTML(200, "open/index.html", gin.H{})
}
