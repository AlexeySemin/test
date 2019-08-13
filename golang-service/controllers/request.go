package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AlexeySemin/test/golang-service/validator"
)

func DecodeAndValidateRequest(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()

	err := validator.Validate(v)
	if err != nil {
		return err
	}

	return nil
}
