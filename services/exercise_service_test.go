package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/fakes"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/mocks"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExerciseCreate(t *testing.T) {
	exerciseRepo := new(mocks.ExerciseRepository)
	userRepo := new(mocks.UserRepository)
	exerciseService := services.NewExerciseService(exerciseRepo, userRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()
		fakeRes := fakes.FakeExerciseCreateResponse()
		fakeUser := fakes.FakeUser()

		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(fakeUser, nil).Once()
		exerciseRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Exercise")).Return(nil).Once()

		res, err := exerciseService.Create(context.TODO(), &fakeReq)

		assert.NoError(t, err)
		assert.Equal(t, res.Exercise.Description, fakeRes.Exercise.Description)
		assert.Equal(t, res.User.ID, fakeRes.User.ID)
		exerciseRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})

	t.Run("fail get user in repository", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()

		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(models.User{}, models.ErrInternalServerError).Once()

		res, err := exerciseService.Create(context.TODO(), &fakeReq)

		assert.Equal(t, models.ErrInternalServerError, err)
		assert.Equal(t, res.Exercise.ID, primitive.NilObjectID)
		assert.Equal(t, res.User.ID, primitive.NilObjectID)
		exerciseRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})

	t.Run("fail user not found in repository", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()

		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(models.User{}, nil).Once()

		res, err := exerciseService.Create(context.TODO(), &fakeReq)

		assert.Equal(t, models.ErrNotFound, err)
		assert.Equal(t, res.Exercise.ID, primitive.NilObjectID)
		assert.Equal(t, res.User.ID, primitive.NilObjectID)
		exerciseRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})

	t.Run("fail date is invalid time", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()
		fakeReq.Date = "2022-02-09"
		fakeUser := fakes.FakeUser()

		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(fakeUser, nil).Once()

		res, err := exerciseService.Create(context.TODO(), &fakeReq)

		assert.Equal(t, models.ErrBadInput, err)
		assert.Equal(t, res.Exercise.ID, primitive.NilObjectID)
		assert.Equal(t, res.User.ID, primitive.NilObjectID)
		exerciseRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})

	t.Run("fail create exercise in repository", func(t *testing.T) {
		fakeReq := fakes.FakeExerciseCreateRequest()
		fakeUser := fakes.FakeUser()

		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(fakeUser, nil).Once()
		exerciseRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Exercise")).Return(models.ErrInternalServerError).Once()

		res, err := exerciseService.Create(context.TODO(), &fakeReq)

		assert.Equal(t, models.ErrInternalServerError, err)
		assert.Equal(t, res.Exercise.ID, primitive.NilObjectID)
		assert.Equal(t, res.User.ID, primitive.NilObjectID)
		exerciseRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})
}

func TestExerciseLogs(t *testing.T) {
	userRepo := new(mocks.UserRepository)
	exerciseRepo := new(mocks.ExerciseRepository)
	exerciseService := services.NewExerciseService(exerciseRepo, userRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		userID := fakes.UserID
		fakeUser := fakes.FakeUser()
		fakeExercises := []models.Exercise{
			fakes.FakeExercise(),
		}

		userRepo.On("GetByID", mock.Anything, userID).Return(fakeUser, nil).Once()
		exerciseRepo.On("GetByUserID", mock.Anything, userID).Return(fakeExercises, nil).Once()

		res, err := exerciseService.GetByUserID(context.TODO(), userID)
		assert.Equal(t, nil, err)
		assert.Equal(t, userID, res.User.ID)
		assert.Equal(t, 1, res.Count)
		assert.Equal(t, fakeExercises[0].ID, res.Exercises[0].ID)
		userRepo.AssertExpectations(t)
		exerciseRepo.AssertExpectations(t)
	})

	t.Run("fail get user in repository", func(t *testing.T) {
		userID := fakes.UserID

		userRepo.On("GetByID", mock.Anything, userID).Return(models.User{}, models.ErrInternalServerError).Once()

		res, err := exerciseService.GetByUserID(context.TODO(), userID)
		assert.Equal(t, models.ErrInternalServerError, err)
		assert.Equal(t, primitive.NilObjectID, res.User.ID)
		assert.Equal(t, 0, res.Count)
		userRepo.AssertExpectations(t)
		exerciseRepo.AssertExpectations(t)
	})

	t.Run("fail user not found in repository", func(t *testing.T) {
		userID := fakes.UserID

		userRepo.On("GetByID", mock.Anything, userID).Return(models.User{}, nil).Once()

		res, err := exerciseService.GetByUserID(context.TODO(), userID)
		assert.Equal(t, models.ErrNotFound, err)
		assert.Equal(t, primitive.NilObjectID, res.User.ID)
		assert.Equal(t, 0, res.Count)
		userRepo.AssertExpectations(t)
		exerciseRepo.AssertExpectations(t)
	})

	t.Run("fail get exercises in repository", func(t *testing.T) {
		userID := fakes.UserID
		fakeUser := fakes.FakeUser()

		userRepo.On("GetByID", mock.Anything, userID).Return(fakeUser, nil).Once()
		exerciseRepo.On("GetByUserID", mock.Anything, userID).Return([]models.Exercise{}, models.ErrInternalServerError).Once()

		res, err := exerciseService.GetByUserID(context.TODO(), userID)
		assert.Equal(t, models.ErrInternalServerError, err)
		assert.Equal(t, primitive.NilObjectID, res.User.ID)
		assert.Equal(t, 0, res.Count)
		userRepo.AssertExpectations(t)
		exerciseRepo.AssertExpectations(t)
	})

	t.Run("fail exercises is not found in repository", func(t *testing.T) {
		userID := fakes.UserID
		fakeUser := fakes.FakeUser()

		userRepo.On("GetByID", mock.Anything, userID).Return(fakeUser, nil).Once()
		exerciseRepo.On("GetByUserID", mock.Anything, userID).Return([]models.Exercise{}, nil).Once()

		res, err := exerciseService.GetByUserID(context.TODO(), userID)
		assert.Equal(t, models.ErrNotFound, err)
		assert.Equal(t, primitive.NilObjectID, res.User.ID)
		assert.Equal(t, 0, res.Count)
		userRepo.AssertExpectations(t)
		exerciseRepo.AssertExpectations(t)
	})
}
