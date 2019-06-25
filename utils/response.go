package utils

import "github.com/kataras/iris"
// SuccessJson 成功返回的json
func SuccessJson(data... interface{}) iris.Map  {
	json := iris.Map{
		"code": 200,
		"status": "success",
	}
	if len(data) > 0 {
		json["data"] = data[0]
	}
	if len(data) > 1 {
		json["message"] = data[1]
	}
	return json
}
// FailureJson 失败时返回的json
func FailureJson(data... interface{}) iris.Map  {
	json := iris.Map{
		"code": 404,
		"status": "failure",
	}
	if len(data) > 1 {
		json["code"] = data[1]
	}
	if len(data) > 0 {
		json["errors"] = data[0]
	}
	return json
}
