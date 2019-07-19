package controllers

import (
	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/jinzhu/gorm"
)

type DBSAController struct {
	db         *gorm.DB
	repository *repositories.DBSARepository
}

func NewDBSAController(db *gorm.DB, repository *repositories.DBSARepository) *DBSAController {
	return &DBSAController{db, repository}
}
