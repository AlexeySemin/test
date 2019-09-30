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

// FillNewsDB godoc
// @Summary Create news
// @Description create news
// @Accept mpfd
// @Produce  json
// @Param Count body request.CreateNews true "Count of news"
// @Success 201 {object} response.LogOnly Created
// @Failure 400 {object} response.Response Bad request
// @Failure 500 {object} response.Response Internal server error
// @Router /api/news [post]
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

// ClearDB godoc
// @Summary Delete news
// @Description delete news
// @Produce  json
// @Success 200 {object} response.LogOnly Deleted
// @Failure 500 {object} response.Response Internal server error
// @Router /api/news [delete]
func (cc *CommonController) ClearDB(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	err := cc.repository.ClearNewsTable()
	if err != nil {
		response.Send(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	end := time.Now()
	resp := response.NewLog(start, end)

	response.Send(w, resp, "News were deleted", http.StatusOK)
}
