package repositories

import (
	"github.com/jinzhu/gorm"
)

type SSARepository struct {
	db *gorm.DB
}

func NewSSARepository(db *gorm.DB) *SSARepository {
	return &SSARepository{db}
}
