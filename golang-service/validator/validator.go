package validator

import (
	v "gopkg.in/go-playground/validator.v9"
)

// Validate validates struct
func Validate(validatedStruct interface{}) error {
	validate := v.New()
	return validate.Struct(validatedStruct)
}
