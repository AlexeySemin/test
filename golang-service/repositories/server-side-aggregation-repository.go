package repositories

import (
	"database/sql"

	"github.com/AlexeySemin/test/golang-service/models"
	"github.com/jinzhu/gorm"
)

type SSARepository struct {
	db *gorm.DB
}

func NewSSARepository(db *gorm.DB) *SSARepository {
	return &SSARepository{db}
}

func (ssar *SSARepository) GetNews() ([]*models.News, error) {
	var news []*models.News

	err := ssar.db.Find(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (ssar *SSARepository) GetRatingsRows() (*sql.Rows, error) {
	rows, err := ssar.db.Raw("SELECT rating FROM news").Rows()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
