package deeplink

import (
	"net/url"

	"zodream.cn/godream/configs"
)

func Encode(path string, params map[string]interface{}) string {
	link := configs.Config.Route.Deeplink + "://" + path
	if len(params) < 1 {
		return link
	}
	var uri url.URL
	q := uri.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	return link + "?" + q.Encode()
}
