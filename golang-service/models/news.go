package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type News struct {
	ID           uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt    time.Time  `gorm:"default:current_timestamp;index"`
	DeletedAt    *time.Time `gorm:"index" json:"-"`
	Title        string     `gorm:"type:varchar(100);not null"`
	Announcement string     `gorm:"type:varchar(255);not null"`
	Text         string     `gorm:"not null"`
	Rating       int        `gorm:"default:0"`
}
