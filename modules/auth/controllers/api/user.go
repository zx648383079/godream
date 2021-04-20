package api

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/auth/rbac"
	"zodream.cn/godream/utils/response"
)

func UserRole(c *gin.Context) {
	api := c.Keys["json"].(response.IJsonResponse)
	role, roles, permissions := rbac.GetUserRole(c.Keys["user_id"].(uint))
	c.JSON(200, api.Render(gin.H{
		"role":        role,
		"roles":       roles,
		"permissions": permissions,
	}))
}
