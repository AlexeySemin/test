package controllers

type createNewsRequest struct {
	Count int `validate:"required,max=1000000"`
}
