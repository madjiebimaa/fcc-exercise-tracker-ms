package mocks

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) Register(ctx context.Context, user *models.User) error {
	ret := u.Called(ctx, user)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = ref(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
