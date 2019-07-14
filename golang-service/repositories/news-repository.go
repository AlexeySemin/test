package repositories

import (
	"github.com/jinzhu/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db}
}
