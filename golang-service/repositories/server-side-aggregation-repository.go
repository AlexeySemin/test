package repositories

import (
	"github.com/jinzhu/gorm"

	"github.com/AlexeySemin/test/golang-service/models"
)

type SSARepository struct {
	db *gorm.DB
}

func NewSSARepository(db *gorm.DB) *SSARepository {
	return &SSARepository{db}
}

func (ssar *SSARepository) GetNews() ([]*models.News, error) {
	var news []*models.News

	// SELECT * FROM "news"  WHERE "news"."deleted_at" IS NULL
	err := ssar.db.Find(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}
