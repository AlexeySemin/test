package repositories

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/AlexeySemin/test/golang-service/db/postgres"
	"github.com/AlexeySemin/test/golang-service/models"
	"github.com/Pallinder/go-randomdata"
	"github.com/jinzhu/gorm"
)

type CommonRepository struct {
	db *gorm.DB
}

func NewCommonRepository(db *gorm.DB) *CommonRepository {
	return &CommonRepository{db}
}

func (cr *CommonRepository) CreateNews(count int) error {
	threadsCount := 20
	part := count
	remainder := 0

	if count < threadsCount {
		threadsCount = 1
	} else {
		remainder = count % threadsCount
		part = count / threadsCount
		if remainder != 0 {
			threadsCount++
		}
	}

	threadsErrors := []string{}
	ch := make(chan error)
	var wg sync.WaitGroup
	wg.Add(threadsCount)

	for i := 1; i <= threadsCount; i++ {
		go func(i int) {
			var news []*models.News

			defer wg.Done()

			if i == threadsCount && remainder != 0 {
				part = remainder
			}

			for j := 1; j <= part; j++ {
				rndDateString := randomdata.FullDateInRange("2019-01-01", "2019-08-22") + " " + strconv.Itoa(randomdata.Number(10, 24)) + ":" + strconv.Itoa(randomdata.Number(10, 60))
				tpl := "Monday _2 Jan 2006 15:04"
				rndDate, _ := time.Parse(tpl, rndDateString)

				oneNews := &models.News{
					CreatedAt:    rndDate,
					Title:        randomdata.State(randomdata.Large),
					Announcement: randomdata.Address(),
					Text:         randomdata.Paragraph(),
					Rating:       randomdata.Number(101),
				}
				news = append(news, oneNews)
			}

			newsToSave := make([]interface{}, len(news))
			for idx, n := range news {
				newsToSave[idx] = n
			}

			_, err := postgres.BatchInsert(cr.db, newsToSave)
			if err != nil {
				ch <- err
				return
			}
		}(i)
	}

	go func() {
		for err := range ch {
			if err != nil {
				threadsErrors = append(threadsErrors, err.Error())
			}
		}
	}()

	wg.Wait()

	if len(threadsErrors) == 0 {
		return nil
	}

	return errors.New(strings.Join(threadsErrors, ", "))
}

func (cr *CommonRepository) ClearDB() error {
	return cr.ClearNewsTable()
}

func (cr *CommonRepository) ClearNewsTable() error {
	var news []*models.News
	return cr.db.Unscoped().Delete(&news).Error
}
