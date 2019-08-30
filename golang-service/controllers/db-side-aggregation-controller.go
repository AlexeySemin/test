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
		response.MinMaxAvgRating
		response.Log
	}{*minMaxAvgResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}

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
		response.PerMonthJSONData
		response.Log
	}{*perMonthJSONResp, *logResp}

	response.Send(w, resp, "", http.StatusOK)
}
