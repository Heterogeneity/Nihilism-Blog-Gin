package database

import "server/global"

// JwtBlacklist 黑名单
type JwtBlacklist struct {
	global.Model
	Jwt string `json:"jwt" gorm:"type:text"` //jwt
}
