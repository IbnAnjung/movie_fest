package utils

import (
	"net/http"

	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type AppError interface {
	Error() string
	ErrorCode() int
}
type ClientError struct {
	Message string
	Code    int
}

func (e ClientError) Error() string {
	return e.Message
}

func (e ClientError) ErrorCode() int {
	return e.Code
}

var (
	DataNotFoundError        = &ClientError{Message: "Data Not Found", Code: http.StatusNotFound}
	DuplicatedDataError      = &ClientError{Message: "Data Already Exists", Code: http.StatusBadRequest}
	UnprocessableEntityError = &ClientError{Message: "Unprocessable", Code: http.StatusUnprocessableEntity}
)

type ValidationError struct {
	Message   string
	Validator enUtil.Validator
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) ErrorCode() int {
	return http.StatusBadRequest
}

type IServerError interface {
	Error() string
	ErrorCode() int
}
type ServerError struct {
}

func (e ServerError) Error() string {
	return "interval server error"
}

func (e ServerError) ErrorCode() int {
	return http.StatusInternalServerError
}
