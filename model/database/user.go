package database

import (
	"github.com/gofrs/uuid"
	"server/global"
	"server/model/appTypes"
)

// User 用户表
type User struct {
	global.Model
	UID       uuid.UUID         `json:"uid" gorm:"type:char(36);unique"`
	Username  string            `json:"username"`
	Password  string            `json:"-"`
	Email     string            `json:"email"`
	Openid    string            `json:"openid"`
	Avatar    string            `json:"avatar" gorm:"size:255"`
	Address   string            `json:"address"`
	Signature string            `json:"signature" gorm:"default:'这个用户很懒，什么都没有写。'"`
	RoleID    appTypes.RoleID   `json:"role_id"`
	Register  appTypes.Register `json:"register"`
	Freeze    bool              `json:"freeze"`
}
