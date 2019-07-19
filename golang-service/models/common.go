package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Common struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"default:current_timestamp;index"`
	DeletedAt *time.Time `gorm:"index"`
}
