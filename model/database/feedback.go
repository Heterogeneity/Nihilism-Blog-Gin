package database

import (
	"github.com/gofrs/uuid"
	"server/global"
)

// Feedback 反馈表
type Feedback struct {
	global.Model
	UserUID uuid.UUID `json:"user_uid" gorm:"type:char(36)"`              //用户UID
	User    User      `json:"-" gorm:"foreignKey:UserUID;references:UID"` //关联用户
	Content string    `json:"content"`                                    //内容
	Reply   string    `json:"reply"`                                      //回复
}
