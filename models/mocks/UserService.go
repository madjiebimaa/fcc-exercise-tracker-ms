package mocks

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/stretchr/testify/mock"
)

type UserService struct {
	mock.Mock
}

func (u *UserService) Register(ctx context.Context, req *requests.UserRegister) (models.User, error) {
	ret := u.Called(ctx, req)

	var r0 models.User
	if ref, ok := ret.Get(0).(func(context.Context, *requests.UserRegister) models.User); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if ref, ok := ret.Get(0).(func(context.Context, *requests.UserRegister) error); ok {
		r1 = ref(ctx, req)
	} else {
		r1 = ret.Error(0)
	}

	return r0, r1
}
