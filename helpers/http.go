package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	case models.ErrBadInput:
		return http.StatusBadRequest
	case models.ErrUnAuthorized:
		return http.StatusUnauthorized
	case models.ErrRedirect:
		return http.StatusTemporaryRedirect
	case models.ErrExpired:
		return http.StatusGone
	default:
		return http.StatusInternalServerError
	}
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, responses.BaseResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func FailResponse(c *gin.Context, status int, errField string, errMessage error) {
	c.JSON(status, responses.BaseResponse{
		Success: false,
		Data:    nil,
		Error: &responses.FieldError{
			Field:   errField,
			Message: errMessage.Error(),
		},
	})
}
