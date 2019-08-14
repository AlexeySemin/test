package controllers

import (
	"net/http"

	"github.com/AlexeySemin/test/golang-service/repositories"
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
	var newsRequest createNewsRequest

	err := DecodeAndValidateRequest(r, &newsRequest)
	if err != nil {
		SendResponse(w, nil, err.Error(), http.StatusBadRequest)
		return
	}

	err = cc.repository.CreateNews(newsRequest.Count)
	if err != nil {
		SendResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	SendResponse(w, nil, "News were created", http.StatusCreated)
}

func (cc *CommonController) ClearDB(w http.ResponseWriter, r *http.Request) {
	err := cc.repository.ClearDB()
	if err != nil {
		SendResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	SendResponse(w, nil, "DB was cleared", http.StatusOK)
}
