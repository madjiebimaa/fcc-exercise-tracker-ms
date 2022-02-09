package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/handlers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/fakes"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := new(mocks.UserService)
	userHandler := handlers.NewUserHandler(userService)

	t.Run("success", func(t *testing.T) {
		sucReq, err := fakes.FakeUserReader("madjiebimaa")
		assert.NoError(t, err)

		userService.On("Register", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil).Times(1)

		req, err := http.NewRequest(http.MethodPost, "/api/users", sucReq)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users", userHandler.Register)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, tt.resBody, rec.Body.Bytes())
		userService.AssertExpectations(t)
	})

	t.Run("fail invalid request", func(t *testing.T) {
		failReq, err := fakes.FakeUserReader("adjie")
		assert.NoError(t, err)
		failBadReqRes, err := json.Marshal(gin.H{
			"message": "user input is not valid",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/api/users", failReq)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users", userHandler.Register)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, failBadReqRes, rec.Body.Bytes())
	})

	t.Run("fail in service", func(t *testing.T) {
		sucReq, err := fakes.FakeUserReader("madjiebimaa")
		assert.NoError(t, err)
		interErr := errors.New("internal server error")
		failInterSerRes, err := json.Marshal(gin.H{
			"message": interErr.Error(),
		})
		assert.NoError(t, err)

		userService.On("Register", mock.Anything, mock.AnythingOfType("*models.User")).Return(interErr).Times(1)

		req, err := http.NewRequest(http.MethodPost, "/api/users", sucReq)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users", userHandler.Register)

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, failInterSerRes, rec.Body.Bytes())
		userService.AssertExpectations(t)
	})
}
