package mocks

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
	"github.com/stretchr/testify/mock"
)

type ExerciseService struct {
	mock.Mock
}

func (e *ExerciseService) Create(ctx context.Context, req *requests.ExerciseCreate) (responses.ExerciseCreate, error) {
	ret := e.Called(ctx, req)

	var r0 responses.ExerciseCreate
	if ref, ok := ret.Get(0).(func(context.Context, *requests.ExerciseCreate) responses.ExerciseCreate); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Get(0).(responses.ExerciseCreate)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, *requests.ExerciseCreate) error); ok {
		r1 = ref(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
