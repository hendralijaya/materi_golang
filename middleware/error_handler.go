package middleware

import (
	"materi-golang/model/web"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	Key string `json:"key,omitempty"`
	Message string `json:"message"`
}

func (e ValidationError) Error(splitedError []string) []ValidationError {
	var errors []ValidationError
	for _, error := range splitedError {
		splittedError := strings.Split(error, "'")
		errors = append(errors, ValidationError{
			Key: splittedError[3],
			Message: "Error :" + splittedError[4] + splittedError[5] + splittedError[6],
		})
	}
	return errors
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	if(c.Errors != nil){
		err := c.Errors.Last()
		if err.Meta == "VALIDATION_ERROR" {
			validationErrors(c, err)
			return
		}
		if err.Meta == "NOT_FOUND" {
			notFoundError(c, err)
			return
		}
		internalServerError(c, err)
	}
}

func validationErrors(c *gin.Context, err *gin.Error){
	splittedError := strings.Split(err.Error(), "\n")
	errors := ValidationError{}.Error(splittedError)
	webResponse := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
		Errors: errors,
		Data:   nil,
	}
	c.JSON(http.StatusBadRequest, webResponse)
}

func notFoundError(c *gin.Context, err *gin.Error){
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Errors: err,
			Data:   nil,
		}
		c.JSON(http.StatusNotFound, webResponse)
}

func internalServerError(c *gin.Context, err *gin.Error) {
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Errors: err,
		Data:   nil,
	}
	c.JSON(http.StatusInternalServerError, webResponse)
}

