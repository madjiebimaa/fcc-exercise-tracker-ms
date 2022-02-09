package fakes

import (
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
)

func FakeExercise() models.Exercise {
	return models.Exercise{
		ID:          exercise_id,
		Description: description,
		Duration:    duration,
		Date:        dateTime,
	}
}

func FakeExerciseCreateRequest() requests.ExerciseCreate {
	return requests.ExerciseCreate{
		UserID:      user_id,
		Description: description,
		Duration:    duration,
		Date:        date,
	}
}

func FakeExerciseCreateResponse() responses.ExerciseCreate {
	return responses.ExerciseCreate{
		User: responses.BaseUser(FakeUser()),
		Exercise: responses.BaseExercise{
			ID:          exercise_id,
			Description: description,
			Duration:    duration,
			Date:        strDate,
		},
	}
}
