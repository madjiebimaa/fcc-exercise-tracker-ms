package fakes

import (
	"encoding/json"
	"strings"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
)

func FakeExercise() models.Exercise {
	return models.Exercise{
		ID:          ExerciseID,
		Description: Description,
		Duration:    Duration,
		Date:        DateTime,
	}
}

func FakeExerciseCreateRequest() requests.ExerciseCreate {
	return requests.ExerciseCreate{
		UserID:      UserID,
		Description: Description,
		Duration:    Duration,
		Date:        Date,
	}
}

func FakeExerciseCreateResponse() responses.ExerciseCreate {
	return responses.ExerciseCreate{
		User: responses.BaseUser(FakeUser()),
		Exercise: responses.BaseExercise{
			ID:          ExerciseID,
			Description: Description,
			Duration:    Duration,
			Date:        StrDate,
		},
	}
}

func FakeExerciseCreateReader() (*strings.Reader, error) {
	reqBody, err := json.Marshal(requests.ExerciseCreate{
		Description: Description,
		Duration:    Duration,
		Date:        Date,
	})
	if err != nil {
		return nil, err
	}
	sucRead := strings.NewReader(string(reqBody))
	return sucRead, nil
}

func FakeExerciseCreateResponseJSON() ([]byte, error) {
	resBody, err := json.Marshal(responses.BaseResponse{
		Success: true,
		Data:    FakeExerciseCreateResponse(),
		Error:   nil,
	})
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
