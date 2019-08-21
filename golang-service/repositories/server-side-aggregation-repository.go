package repositories

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type SSARepository struct {
	db *gorm.DB
}

func NewSSARepository(db *gorm.DB) *SSARepository {
	return &SSARepository{db}
}

func (ssar *SSARepository) GetRatings() (*sql.Rows, error) {
	rows, err := ssar.db.Raw("SELECT rating FROM news").Rows()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
