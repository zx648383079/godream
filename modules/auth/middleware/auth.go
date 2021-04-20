package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authorize(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		val := session.Get("user_id")
		var userId uint
		if val != nil {
			userId = val.(uint)
		}
		if userId < 1 && required {
			c.AbortWithStatus(401)
			return
		}
		c.Keys["user_id"] = userId
		c.Next()
	}
}
