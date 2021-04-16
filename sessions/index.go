package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	cookieNameForSessionID = "zodream"
)

func New() gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	return sessions.Sessions(cookieNameForSessionID, store)
}
