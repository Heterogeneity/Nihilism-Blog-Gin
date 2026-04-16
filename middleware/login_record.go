package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
	"go.uber.org/zap"
	"server/global"
	"server/model/database"
	"server/service"
)

// LoginRecord 是一个中间件，用于记录登录日志
func LoginRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 异步记录日志
		go func() {
			gaodeService := service.ServiceGroupApp.GaodeService
			var userID uint
			var address string
			ip := c.ClientIP()
			loginMethod := c.DefaultQuery("flag", "email")
			userAgent := c.Request.UserAgent()

			if value, exists := c.Get("user_id"); exists {
				if id, ok := value.(uint); ok {
					userID = id
				}
			}

			address = getAddressFromIP(ip, gaodeService)

			os, deviceInfo, browserInfo := parseUserAgent(userAgent)
			login := database.Login{
				UserID:      userID,
				LoginMethod: loginMethod,
				IP:          ip,
				Address:     address,
				OS:          os,
				DeviceInfo:  deviceInfo,
				BrowserInfo: browserInfo,
				Status:      c.Writer.Status(),
			}
			if err := global.DB.Create(&login).Error; err != nil {
				global.Log.Error("登录记录失败", zap.Error(err))
			}
		}()
	}

}

// 获取IP地址对应的地理位置信息
func getAddressFromIP(ip string, gaodeService service.GaodeService) string {
	res, err := gaodeService.GetLocationByIP(ip)
	if err != nil || res.Province == "" {
		return "未知"
	}
	if res.City != "" && res.Province != res.City {
		return res.Province + "-" + res.City
	}
	return res.Province
}

// 解析用户代理（User-Agent）字符串，提取操作系统、设备信息和浏览器信息
func parseUserAgent(userAgent string) (os, deviceInfo, browserInfo string) {
	os = userAgent
	deviceInfo = userAgent
	browserInfo = userAgent
	parser := uaparser.NewFromSaved()
	cli := parser.Parse(userAgent)
	os = cli.Os.Family
	deviceInfo = cli.Device.Family
	browserInfo = cli.UserAgent.Family
	return
}
