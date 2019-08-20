package repositories

import (
	"github.com/jinzhu/gorm"

	"github.com/AlexeySemin/test/golang-service/response"
)

type DBSARepository struct {
	db *gorm.DB
}

func NewDBSARepository(db *gorm.DB) *DBSARepository {
	return &DBSARepository{db}
}

func (dbsar *DBSARepository) GetMinMaxAvgRating() (*response.MinMaxAvgRating, error) {
	var resp response.MinMaxAvgRating

	// SELECT min(rating), max(rating), avg(rating) FROM "news"
	err := dbsar.db.Table("news").
		Select("min(rating), max(rating), avg(rating)").
		Scan(&resp).
		Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
