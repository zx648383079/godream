package platform

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/utils"
)

// PlatformResponse json响应
type PlatformResponse struct {
	Platform *Platform
}

// Render 响应
func (PlatformResponse) Render(data gin.H) gin.H {
	return data
}

// RenderData 成功返回的json data, message
func (r PlatformResponse) RenderData(data ...interface{}) gin.H {
	json := gin.H{}
	if len(data) > 0 {
		json["data"] = data[0]
	}
	if len(data) > 1 {
		json["message"] = data[1]
	}
	return r.Render(json)
}

// RenderPage 响应分页
func (r PlatformResponse) RenderPage(data interface{}, page *utils.Pager) gin.H {
	json := gin.H{
		"data": data,
		"paging": gin.H{
			"limit":  page.Size,
			"offset": page.Current,
			"total":  page.Total,
			"more":   page.IsNext,
		},
	}
	return r.Render(json)
}

// RenderFailure 失败时返回的json message code
func (r PlatformResponse) RenderFailure(data ...interface{}) gin.H {
	json := gin.H{
		"code":    404,
		"message": "error",
	}
	if len(data) > 1 {
		json["code"] = data[1]
	}
	if len(data) > 0 {
		json["message"] = data[0]
	}
	return r.Render(json)
}
