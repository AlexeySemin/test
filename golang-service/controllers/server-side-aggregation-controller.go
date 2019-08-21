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

	rows, err := ssac.repository.GetRatings()
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
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

	minMaxAvgResp := &response.MinMaxAvgRating{
		Min: min,
		Max: max,
		Avg: avg,
	}

	end := time.Now()
	logResp := response.NewLog(start, end)
	resp := struct {
		response.MinMaxAvgRating
		response.Log
	}{*minMaxAvgResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}
