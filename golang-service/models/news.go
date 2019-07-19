package models

type News struct {
	Common
	Title        string `gorm:"type:varchar(100);not null"`
	Announcement string `gorm:"type:varchar(255);not null"`
	Text         string `gorm:"type:varchar(255);not null"`
	Rating       *int   `gorm:"default:0"`
}
