package mocks

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseRepository struct {
	mock.Mock
}

func (e *ExerciseRepository) Create(ctx context.Context, exercise *models.Exercise) error {
	ret := e.Called(ctx, exercise)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *models.Exercise) error); ok {
		r0 = ref(ctx, exercise)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (e *ExerciseRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Exercise, error) {
	ret := e.Called(ctx, userID)

	var r0 []models.Exercise
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) []models.Exercise); ok {
		r0 = ref(ctx, userID)
	} else {
		r0 = ret.Get(0).([]models.Exercise)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
