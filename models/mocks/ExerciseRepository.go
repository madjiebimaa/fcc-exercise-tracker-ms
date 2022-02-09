package mocks

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/stretchr/testify/mock"
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
