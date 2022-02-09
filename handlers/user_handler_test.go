package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/handlers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
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
		fakeReq := fakes.FakeUserRegisterRequest()
		fakeRes, err := fakes.FakeUserJSON()
		assert.NoError(t, err)
		fakeReqReader, err := fakes.FakeUserRegisterReader()
		assert.NoError(t, err)
		fakeUser := fakes.FakeUser()

		userService.On("Register", mock.Anything, &fakeReq).Return(fakeUser, nil).Once()

		req, err := http.NewRequest(http.MethodPost, "/api/users", fakeReqReader)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users", userHandler.Register)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, fakeRes, rec.Body.Bytes())
		userService.AssertExpectations(t)
	})

	t.Run("fail request body is not valid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/users", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users", userHandler.Register)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		// TODO: create fakes data for error handler
		userService.AssertExpectations(t)
	})

	t.Run("fail register user in service", func(t *testing.T) {
		fakeReq := fakes.FakeUserRegisterRequest()
		fakeReqReader, err := fakes.FakeUserRegisterReader()
		assert.NoError(t, err)

		userService.On("Register", mock.Anything, &fakeReq).Return(models.User{}, models.ErrInternalServerError).Once()

		req, err := http.NewRequest(http.MethodPost, "/api/users", fakeReqReader)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users", userHandler.Register)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		// TODO: create fakes data for error handler
		userService.AssertExpectations(t)
	})
}
