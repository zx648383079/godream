package platform

import (
	"zodream/utils"

	"github.com/kataras/iris/v12"
)

// PlatformResponse json响应
type PlatformResponse struct {
	Platform *Platform
}

// Render 响应
func (PlatformResponse) Render(data iris.Map) iris.Map {
	return data
}

// RenderData 成功返回的json data, message
func (r PlatformResponse) RenderData(data ...interface{}) iris.Map {
	json := iris.Map{}
	if len(data) > 0 {
		json["data"] = data[0]
	}
	if len(data) > 1 {
		json["message"] = data[1]
	}
	return r.Render(json)
}

// RenderPage 响应分页
func (r PlatformResponse) RenderPage(data []interface{}, page utils.Pager) iris.Map {
	json := iris.Map{
		"data": data,
		"paging": iris.Map{
			"limit":  page.Size,
			"offset": page.Current,
			"total":  page.Total,
			"more":   page.IsNext,
		},
	}
	return r.Render(json)
}

// RenderFailure 失败时返回的json message code
func (r PlatformResponse) RenderFailure(data ...interface{}) iris.Map {
	json := iris.Map{
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
