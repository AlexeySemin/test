package controllers

import (
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

	news, err := ssac.repository.GetNews()
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	min := 0
	max := 0
	sumRating := 0
	count := len(news)

	for _, n := range news {
		if n.Rating < min {
			min = n.Rating
		}
		if n.Rating > max {
			max = n.Rating
		}
		sumRating += n.Rating
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
		response.LogResponse
	}{*minMaxAvgResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}
