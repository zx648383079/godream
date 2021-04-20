package deeplink

import (
	"net/url"
	"strconv"
	"strings"

	"zodream.cn/godream/configs"
)

func Decode(link string) string {
	if link == "" || link[:1] == "#" || link[:11] == "javascript:" {
		return link
	}
	builder, err := url.Parse(link)
	if err != nil {
		return link
	}
	if builder.Scheme == "" || builder.Scheme != configs.Config.Route.Deeplink {
		return link
	}
	if builder.Host == "" {
		return ""
	}
	if builder.Host == "chat" {
		return "/chat"
	}
	if builder.Path == "" || builder.Path == "/" {
		return "/"
	}
	backendMap := []string{"b", "admin", "backend", "system"}
	isBackend := false
	for _, v := range backendMap {
		if v == builder.Host {
			isBackend = true
			break
		}
	}
	items := strings.Split(strings.Trim(builder.Path, "/"), "/")
	path := items[0]
	if isBackend && path == "friend_link" {
		return "/contact/admin/friend_link"
	}
	params := items[1:]
	id, _ := strconv.Atoi(params[0])
	if isBackend && path == "user" && id > 0 {
		return "/auth/admin/user/edit?id=" + params[0]
	}
	if isBackend && path == "order" && id > 0 {
		return "/shop/admin/order/info?id=" + params[0]
	}
	return ""
}
