package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/service"
	"server/utils"
	"strconv"
)

var jwtService = service.ServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求中的Access Token和Refresh Token
		accessToken := utils.GetAccessToken(c)
		refreshToken := utils.GetRefreshToken(c)

		// 检查Refresh Token是否在黑名单中，如果是，则清除Refresh Token并返回未授权错误
		if jwtService.IsInBlacklist(refreshToken) {
			utils.ClearRefreshToken(c)
			response.NoAuth("黑名单token无效", c)
			c.Abort() // 终止请求的后续处理
			return
		}
		// 创建一个JWT实例，用于后续的token解析与验证
		j := utils.NewJWT()
		// 解析Access Token
		claims, err := j.ParseAccessToken(accessToken)
		if err != nil {
			// 如果解析失败并且Access Token为空或过期
			if accessToken == "" || errors.Is(err, utils.TokenExpired) {
				// 尝试解析Refresh Token
				refreshClaims, err := j.ParseRefreshToken(refreshToken)
				if err != nil {
					// 如果Refresh Token也无法解析，清除Refresh Token并返回未授权错误
					utils.ClearRefreshToken(c)
					response.NoAuth("refresh-token过期或无效", c)
					c.Abort()
					return
				}

				// 如果Refresh Token有效，通过其UserID获取用户信息
				var user database.User
				if err := global.DB.Select("uid", "role_id").Take(&user, refreshClaims.UserID).Error; err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("用户不存在", c)
					c.Abort()
					return
				}
				newAccessClaims := j.CreateAccessClaims(request.BaseClaims{
					UserID: refreshClaims.UserID,
					UID:    user.UID,
					RoleID: user.RoleID,
				})
				newAccessToken, err := j.CreateAccessToken(newAccessClaims)
				if err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("创建新token失败", c)
					c.Abort()
					return
				}
				c.Header("new-access-token", newAccessToken)
				c.Header("new-access-expires-at", strconv.FormatInt(newAccessClaims.ExpiresAt.Unix(), 10))
				c.Set("claims", &newAccessClaims)
				c.Next()
				return
			}
			utils.ClearRefreshToken(c)
			response.NoAuth("无效的access-token", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
