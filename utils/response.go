package utils

import (
	"log"
	"net/http"

	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type metaPaginationResponse struct {
	Limit     int   `json:"limit"`
	Page      int   `json:"page"`
	TotalPage int   `json:"total_page"`
	TotalRaw  int64 `json:"total_raw"`
}

type paginationResponse struct {
	Message string                 `json:"message"`
	Meta    metaPaginationResponse `json:"meta"`
	Data    interface{}            `json:"data"`
}

type validationResponse struct {
	response
	ValidationErrors []string `json:"validation_errors"`
}

func SuccessResponse(c *gin.Context, httpStatusCode int, message string, data interface{}) {
	c.JSON(httpStatusCode, response{
		Message: message,
		Data:    data,
	})
}

func PaginationResponse(c *gin.Context, httpStatusCode int, message string, meta enUtil.MetaPagination, data interface{}) {
	c.JSON(httpStatusCode, paginationResponse{
		Message: message,
		Meta: metaPaginationResponse{
			Limit:     meta.Limit,
			Page:      meta.Page,
			TotalPage: meta.TotalPage,
			TotalRaw:  meta.TotalRaw,
		},
		Data: data,
	})
}

func ErrorValidationResponse(c *gin.Context, message string, errors []string) {
	c.JSON(http.StatusBadRequest, validationResponse{
		response: response{
			Message: message,
		},
		ValidationErrors: errors,
	})
}

func ErrorResponse(c *gin.Context, httpStatusCode int, message string) {
	c.JSON(httpStatusCode, response{
		Message: message,
	})
}

func GeneralErrorResponse(c *gin.Context, err error) {
	if err, ok := err.(ValidationError); ok {
		validationErrors := err.Validator.GetValidationErrors()
		ErrorValidationResponse(c, err.Error(), validationErrors)
		return
	}

	if err, ok := err.(*ClientError); ok {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("Internal Server Error: %s", err.Error())
	ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
}
