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

func TestUserRegister(t *testing.T) {
	userRepo := new(mocks.UserRepository)
	userService := services.NewUserService(userRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		fakeReq := fakes.FakeUserRegisterRequest()
		fakeUser := fakes.FakeUser()

		userRepo.On("Register", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil).Once()

		user, err := userService.Register(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		assert.Equal(t, fakeUser.UserName, user.UserName)
		userRepo.AssertExpectations(t)
	})

	t.Run("fail register user in repository", func(t *testing.T) {
		fakeReq := fakes.FakeUserRegisterRequest()

		userRepo.On("Register", mock.Anything, mock.AnythingOfType("*models.User")).Return(models.ErrInternalServerError).Once()

		user, err := userService.Register(context.TODO(), &fakeReq)
		assert.Equal(t, models.ErrInternalServerError, err)
		assert.Equal(t, primitive.NilObjectID, user.ID)
		userRepo.AssertExpectations(t)
	})
}
