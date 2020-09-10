package response

import (
	"zodream/utils"

	"github.com/kataras/iris/v12"
)

// IJsonResponse json响应接口
type IJsonResponse interface {
	Render(data iris.Map) iris.Map
	RenderData(data ...interface{}) iris.Map
	RenderPage(data []interface{}, page utils.Pager) iris.Map
	RenderFailure(data ...interface{}) iris.Map
}

// JSONResponse json响应
type JSONResponse struct {
}

// Render 响应
func (JSONResponse) Render(data iris.Map) iris.Map {
	return data
}

// RenderData 成功返回的json data, message
func (r JSONResponse) RenderData(data ...interface{}) iris.Map {
	json := iris.Map{
		"code":   200,
		"status": "success",
	}
	if len(data) > 0 {
		json["data"] = data[0]
	}
	if len(data) > 1 {
		json["message"] = data[1]
	}
	return r.Render(json)
}

// RenderPage 响应分页
func (r JSONResponse) RenderPage(data []interface{}, page utils.Pager) iris.Map {
	json := iris.Map{
		"code":   200,
		"status": "success",
		"data":   data,
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
func (r JSONResponse) RenderFailure(data ...interface{}) iris.Map {
	json := iris.Map{
		"code":   404,
		"status": "failure",
	}
	if len(data) > 1 {
		json["code"] = data[1]
	}
	if len(data) > 0 {
		json["message"] = data[0]
	}
	return r.Render(json)
}
