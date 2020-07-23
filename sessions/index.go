package sessions

import (
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "zodream"
	// Driver session 使用
	Driver = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
