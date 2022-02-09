package mocks

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u *UserRepository) GetByID(ctx context.Context, userID primitive.ObjectID) (models.User, error) {
	ret := u.Called(ctx, userID)

	var r0 models.User
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) models.User); ok {
		r0 = ref(ctx, userID)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
