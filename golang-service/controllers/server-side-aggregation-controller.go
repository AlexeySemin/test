package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/AlexeySemin/test/golang-service/response"
	"github.com/jinzhu/gorm"
)

type SSAController struct {
	db         *gorm.DB
	repository *repositories.SSARepository
}

func NewSSAController(db *gorm.DB) *SSAController {
	repository := repositories.NewSSARepository(db)
	return &SSAController{db, repository}
}

func (ssac *SSAController) GetMinMaxAvgRating(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	useRows := r.FormValue("use_rows")
	var minMaxAvgResp *response.MinMaxAvgRating
	var err error

	if useRows == "" || useRows == "false" {
		minMaxAvgResp, err = ssac.getMinMaxAvgRating()
	} else {
		minMaxAvgResp, err = ssac.getMinMaxAvgRatingUsingRows()
	}

	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	logResp := response.NewLog(start, end)
	resp := struct {
		response.MinMaxAvgRating
		response.Log
	}{*minMaxAvgResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}

func (ssac *SSAController) getMinMaxAvgRating() (*response.MinMaxAvgRating, error) {
	news, err := ssac.repository.GetNews()
	if err != nil {
		return nil, err
	}

	min := 0
	max := 0
	sumRating := 0
	count := 0

	for _, oneNews := range news {
		if oneNews.Rating < min {
			min = oneNews.Rating
		}
		if oneNews.Rating > max {
			max = oneNews.Rating
		}
		sumRating += oneNews.Rating
		count++
	}

	avg := float64(sumRating) / float64(count)

	return &response.MinMaxAvgRating{
		Min: min,
		Max: max,
		Avg: avg,
	}, nil
}

func (ssac *SSAController) getMinMaxAvgRatingUsingRows() (*response.MinMaxAvgRating, error) {
	rows, err := ssac.repository.GetRatingsRows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	min := 0
	max := 0
	sumRating := 0
	count := 0
	var rating int

	for rows.Next() {
		if err := rows.Scan(&rating); err != nil {
			log.Print(err)
		}
		if rating < min {
			min = rating
		}
		if rating > max {
			max = rating
		}
		sumRating += rating
		count++
	}

	avg := float64(sumRating) / float64(count)

	return &response.MinMaxAvgRating{
		Min: min,
		Max: max,
		Avg: avg,
	}, nil
}
