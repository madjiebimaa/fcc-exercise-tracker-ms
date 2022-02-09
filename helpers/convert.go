package helpers

import (
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
)

func ToBaseExercise(exercise models.Exercise) responses.BaseExercise {
	strDate := exercise.Date.Format(time.RFC1123)
	return responses.BaseExercise{
		ID:          exercise.ID,
		Description: exercise.Description,
		Duration:    exercise.Duration,
		Date:        strDate,
	}
}

func ToBaseExercises(exercises []models.Exercise) []responses.BaseExercise {
	var baseExercises []responses.BaseExercise
	for _, exercise := range exercises {
		baseExercises = append(baseExercises, ToBaseExercise(exercise))
	}
	return baseExercises
}
