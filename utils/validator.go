package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Err       error
	validator validator.Validate
}

func NewValidator() (Validator, error) {
	v := validator.New()

	return Validator{
		validator: *v,
	}, nil
}

func (v Validator) ValidateStruct(input interface{}) error {
	v.Err = v.validator.Struct(input)

	if v.Err == nil {
		return nil
	}

	return ValidationError{
		Message:   "invalid payload",
		Validator: v,
	}
}

func (v Validator) GetValidationErrors() []string {
	errors := []string{}

	for _, e := range v.Err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("%s %s", e.Field(), e.Error()))
	}

	return errors
}
