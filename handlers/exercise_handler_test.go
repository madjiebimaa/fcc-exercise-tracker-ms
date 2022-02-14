package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/handlers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/fakes"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/mocks"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExerciseCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	exerciseService := new(mocks.ExerciseService)
	exerciseHandler := handlers.NewExerciseHandler(exerciseService)

	t.Run("success", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()
		fakeRes := fakes.FakeExerciseCreateResponse()
		fakeResJSON, err := fakes.FakeExerciseCreateResponseJSON()
		assert.NoError(t, err)
		fakeStrRead, err := fakes.FakeExerciseCreateReader()
		assert.NoError(t, err)

		exerciseService.On("Create", mock.Anything, &fakeReq).Return(fakeRes, nil).Once()

		url := fmt.Sprintf("/api/users/%s/exercises", fakes.UserIDStr)
		req, err := http.NewRequest(http.MethodPost, url, fakeStrRead)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/exercises", exerciseHandler.Create)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, fakeResJSON, rec.Body.Bytes())
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail user id is not found in url param", func(t *testing.T) {
		fakeStrRead, err := fakes.FakeExerciseCreateReader()
		assert.NoError(t, err)

		url := fmt.Sprintf("/api/users/%s/exercises", "")
		req, err := http.NewRequest(http.MethodPost, url, fakeStrRead)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/exercises", exerciseHandler.Create)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		// TODO: create fakes data for error handler
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail request body is not valid", func(t *testing.T) {
		url := fmt.Sprintf("/api/users/%s/exercises", fakes.UserIDStr)
		req, err := http.NewRequest(http.MethodPost, url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/exercises", exerciseHandler.Create)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		// TODO: create fakes data for error handler
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail url param is not valid id", func(t *testing.T) {
		fakeStrRead, err := fakes.FakeExerciseCreateReader()
		assert.NoError(t, err)

		url := fmt.Sprintf("/api/users/%s/exercises", "invalid")
		req, err := http.NewRequest(http.MethodPost, url, fakeStrRead)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/exercises", exerciseHandler.Create)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		fmt.Println(rec.Body)
		// TODO: create fakes data for error handler
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail create exercise in service", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()
		fakeStrRead, err := fakes.FakeExerciseCreateReader()
		assert.NoError(t, err)

		exerciseService.On("Create", mock.Anything, &fakeReq).Return(responses.ExerciseCreate{}, models.ErrInternalServerError).Once()

		url := fmt.Sprintf("/api/users/%s/exercises", fakes.UserIDStr)
		req, err := http.NewRequest(http.MethodPost, url, fakeStrRead)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/exercises", exerciseHandler.Create)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		// TODO: create fakes data for error handler
		exerciseService.AssertExpectations(t)
	})
}

func TestExerciseGetByUserID(t *testing.T) {
	exerciseService := new(mocks.ExerciseService)
	exerciseHandler := handlers.NewExerciseHandler(exerciseService)

	t.Run("success", func(t *testing.T) {
		fakeRes := fakes.FakeExerciseLogsResponse()
		fakeResJSON, err := fakes.FakeExerciseLogsResponseJSON()
		assert.NoError(t, err)

		exerciseService.On("GetByUserID", mock.Anything, fakes.UserID).Return(fakeRes, nil).Once()

		url := fmt.Sprintf("/api/users/%s/logs", fakes.UserIDStr)
		req, err := http.NewRequest(http.MethodPost, url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/logs", exerciseHandler.GetByUserID)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, fakeResJSON, rec.Body.Bytes())
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail user id is not found in url param", func(t *testing.T) {
		url := fmt.Sprintf("/api/users/%s/logs", "")
		req, err := http.NewRequest(http.MethodPost, url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/logs", exerciseHandler.GetByUserID)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail url param is not valid id", func(t *testing.T) {
		url := fmt.Sprintf("/api/users/%s/logs", "invalid")
		req, err := http.NewRequest(http.MethodPost, url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/logs", exerciseHandler.GetByUserID)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		exerciseService.AssertExpectations(t)
	})

	t.Run("fail get logs by user id in service", func(t *testing.T) {
		exerciseService.On("GetByUserID", mock.Anything, fakes.UserID).Return(responses.ExerciseLogs{}, models.ErrInternalServerError).Once()

		url := fmt.Sprintf("/api/users/%s/logs", fakes.UserIDStr)
		req, err := http.NewRequest(http.MethodPost, url, nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		r.POST("/api/users/:userID/logs", exerciseHandler.GetByUserID)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		exerciseService.AssertExpectations(t)
	})
}
