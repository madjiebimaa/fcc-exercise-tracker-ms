package services_test

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
// 	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/fakes"
// 	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models/mocks"
// 	"github.com/madjiebimaa/fcc-exercise-tracker-ms/services"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestUserRegister(t *testing.T) {
// 	u := fakes.FakeUser()
// 	interErr := errors.New("internal server error")

// 	testCases := []struct {
// 		name    string
// 		resRepo error
// 		reqSer  models.User
// 		resSer  error
// 	}{
// 		{
// 			name:    "success",
// 			resRepo: nil,
// 			reqSer:  u,
// 			resSer:  nil,
// 		},
// 		{
// 			name:    "fail register in repository",
// 			resRepo: interErr,
// 			reqSer:  u,
// 			resSer:  interErr,
// 		},
// 	}

// 	userRepo := new(mocks.UserRepository)
// 	userService := services.NewUserService(userRepo, 2*time.Second)

// 	for _, tt := range testCases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			userRepo.On("Register", mock.Anything, &u).Return(tt.resRepo).Times(1)

// 			err := userService.Register(context.TODO(), &u)

// 			assert.Equal(t, err, tt.resSer)
// 			userRepo.AssertExpectations(t)
// 		})
// 	}
// }
