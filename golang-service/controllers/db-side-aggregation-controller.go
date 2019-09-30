package controllers

import (
	"net/http"
	"time"

	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/AlexeySemin/test/golang-service/response"
	"github.com/jinzhu/gorm"
)

type DBSAController struct {
	db         *gorm.DB
	repository *repositories.DBSARepository
}

func NewDBSAController(db *gorm.DB) *DBSAController {
	repository := repositories.NewDBSARepository(db)
	return &DBSAController{db, repository}
}

// GetMinMaxAvgRating godoc
// @Summary DB side aggregation of the min, max, avg news rating
// @Description get min, max, avg news rating
// @Produce json
// @Success 200 {object} response.MinMaxAvgRating
// @Failure 500 {object} response.Response Internal server error
// @Router /api/dbsa/news/min-max-avg-rating [get]
func (dbsac *DBSAController) GetMinMaxAvgRating(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	minMaxAvgResp, err := dbsac.repository.GetMinMaxAvgRating()
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	logResp := response.NewLog(start, end)
	resp := struct {
		response.MinMaxAvg
		response.Log
	}{*minMaxAvgResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}

// GetPerMonthJSONData godoc
// @Summary DB side aggregation of the min, max, avg, count news per month
// @Description get min, max, avg, count news per month
// @Produce json
// @Success 200 {object} response.PerMonthJSONData
// @Failure 500 {object} response.Response Internal server error
// @Router /api/dbsa/news/per-month-json-data [get]
func (dbsac *DBSAController) GetPerMonthJSONData(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	perMonthJSONResp, err := dbsac.repository.GetPerMonthJSONData()
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	logResp := response.NewLog(start, end)
	resp := struct {
		response.PerMonthJSON
		response.Log
	}{*perMonthJSONResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}
