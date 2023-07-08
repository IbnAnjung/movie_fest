package utils

//go:generate mockery --name Validator
type Validator interface {
	ValidateStruct(input interface{}) error
	GetValidationErrors() []string
}
