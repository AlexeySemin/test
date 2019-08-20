package controllers

import (
	"net/http"
	"time"

	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/AlexeySemin/test/golang-service/request"
	"github.com/AlexeySemin/test/golang-service/response"
	"github.com/jinzhu/gorm"
)

type CommonController struct {
	db         *gorm.DB
	repository *repositories.CommonRepository
}

func NewCommonController(db *gorm.DB) *CommonController {
	repository := repositories.NewCommonRepository(db)
	return &CommonController{db, repository}
}

func (cc *CommonController) FillNewsDB(w http.ResponseWriter, r *http.Request) {
	var newsRequest request.CreateNews

	err := request.DecodeAndValidate(r, &newsRequest)
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	start := time.Now()

	err = cc.repository.CreateNews(newsRequest.Count)
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	resp := response.NewLog(start, end)

	response.Send(w, resp, "News were created", http.StatusCreated)
}

func (cc *CommonController) ClearDB(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	err := cc.repository.ClearDB()
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	resp := response.NewLog(start, end)

	response.Send(w, resp, "DB was cleared", http.StatusOK)
}
