package database

import "server/global"

// Advertisement 广告表
type Advertisement struct {
	global.Model
	AdImage string `json:"ad_image" gorm:"size:255"`
	Image   Image  `json:"-" gorm:"foreignKey:AdImage;references:URL"`
	Link    string `json:"link"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
