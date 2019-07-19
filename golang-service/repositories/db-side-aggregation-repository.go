package repositories

import (
	"github.com/jinzhu/gorm"
)

type DBSARepository struct {
	db *gorm.DB
}

func NewDBSARepository(db *gorm.DB) *DBSARepository {
	return &DBSARepository{db}
}
