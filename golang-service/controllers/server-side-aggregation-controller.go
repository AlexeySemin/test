package controllers

import (
	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/jinzhu/gorm"
)

type SSAController struct {
	db         *gorm.DB
	repository *repositories.SSARepository
}

func NewSSAController(db *gorm.DB, repository *repositories.SSARepository) *SSAController {
	return &SSAController{db, repository}
}
