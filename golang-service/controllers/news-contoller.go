package controllers

import (
	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/jinzhu/gorm"
)

type NewsController struct {
	db         *gorm.DB
	repository *repositories.NewsRepository
}

func NewNewsController(db *gorm.DB, repository *repositories.NewsRepository) *NewsController {
	return &NewsController{db, repository}
}
