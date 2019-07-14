package controllers

import (
	"github.com/AlexeySemin/test/golang-service/repositories"
	"github.com/jinzhu/gorm"
)

type ProductController struct {
	db         *gorm.DB
	repository *repositories.ProductRepository
}

func NewProductController(db *gorm.DB, repository *repositories.ProductRepository) *ProductController {
	return &ProductController{db, repository}
}
