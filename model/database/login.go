package database

import "server/global"

// Login 登录表
type Login struct {
	global.Model
	UserID      uint   `json:"user_id"` //用户ID
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	LoginMethod string `json:"login_method"` //登录方式
	IP          string `json:"ip"`           //IP地址
	Address     string `json:"address"`      //登录地址
	OS          string `json:"os"`           //操作系统
	DeviceInfo  string `json:"device_info"`  //设备信息
	BrowserInfo string `json:"browser_info"` //浏览器信息
	Status      int    `json:"status"`       //登录状态
}
