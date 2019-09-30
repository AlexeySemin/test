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

func (dbsar *DBSARepository) GetMinMaxAvgRating() (*response.MinMaxAvg, error) {
	var resp response.MinMaxAvg

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

func (dbsar *DBSARepository) GetPerMonthJSONData() (*response.PerMonthJSON, error) {
	var resp response.PerMonthJSON

	err := dbsar.db.Raw(`
		select (
			'[' || string_agg('{"date":' || '"' || date || '","avgRating":' || avg_rating || ',"minRating":' || min_rating || ',"maxRating":' || max_rating || ',"countNews":' || count_news || '}' ,',') || ']'
		)::json as data
		from (
			select
				date_trunc('month', d.day)::date as date,
				avg(n.rating) as avg_rating,
				min(n.rating) as min_rating,
				max(n.rating) as max_rating,
				count(n.*) as count_news
			from (
				select
					generate_series(date_trunc('year', '2019-01-01'::date), date_trunc('year', '2019-01-01'::date) + '1 year - 1 day'::interval, '1 day'::interval)::date
				) d(day)
			left join news as n on date(n.created_at) = d.day
			where n.deleted_at is null
			group by 1
			order by 1
		) as perMonth
	`).Scan(&resp).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
