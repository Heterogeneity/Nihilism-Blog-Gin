package middleware

import (
	"github.com/gin-gonic/gin"
	"server/model/appTypes"
	"server/model/response"
	"server/utils"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := utils.GetRoleID(c)
		if roleID != appTypes.Admin {
			response.Forbidden("并非管理员", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
