package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"net"
	"server/global"
	"server/model/appTypes"
	"server/model/request"
)

func SetRefreshToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", token, maxAge, host)
}

func ClearRefreshToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", "", -1, host)
}

func setCookie(c *gin.Context, name string, value string, maxAge int, host string) {
	if net.ParseIP(host) != nil {
		c.SetCookie(name, value, maxAge, "/", "", false, true)
	} else {
		c.SetCookie(name, value, maxAge, "/", host, false, true)
	}
}
func GetAccessToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-access-token")
	return token
}

func GetRefreshToken(c *gin.Context) string {
	token, _ := c.Cookie("x-refresh-token")
	return token
}

func GetClaims(c *gin.Context) (*request.JwtCustomClaims, error) {
	token := GetAccessToken(c)
	j := NewJWT()
	claims, err := j.ParseAccessToken(token)
	if err != nil {
		global.Log.Error("access-token获取失败", zap.Error(err))
	}
	return claims, err
}

func GetRefreshClaims(c *gin.Context) (*request.JwtCustomRefreshClaims, error) {
	token := GetRefreshToken(c)
	j := NewJWT()
	claims, err := j.ParseRefreshToken(token)
	if err != nil {
		global.Log.Error("refresh-token获取失败", zap.Error(err))
	}
	return claims, err
}

func GetUserInfo(c *gin.Context) *request.JwtCustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse
	}
}

func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.UserID
	}
}

func GetUUID(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UID
		}
	} else {
		waiUse := claims.(*request.JwtCustomClaims)
		return waiUse.UID
	}
}

func GetRoleID(c *gin.Context) appTypes.RoleID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleID
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.RoleID
	}
}
