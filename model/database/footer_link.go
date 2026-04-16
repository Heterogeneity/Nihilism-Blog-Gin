package database

// FooterLink 页脚链接表
type FooterLink struct {
	Title string `json:"title" gorm:"primaryKey"`
	Link  string `json:"link"`
}
