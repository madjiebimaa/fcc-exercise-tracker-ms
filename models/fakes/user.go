package fakes

import (
	"encoding/json"
	"strings"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
)

func FakeUser() models.User {
	return models.User{
		ID:       UserID,
		UserName: UserName,
	}
}

func FakeUserRegisterRequest() requests.UserRegister {
	return requests.UserRegister{
		UserName: UserName,
	}
}

func FakeUserRegisterReader() (*strings.Reader, error) {
	reqBody, err := json.Marshal(requests.UserRegister{
		UserName: UserName,
	})
	if err != nil {
		return nil, err
	}
	sucRead := strings.NewReader(string(reqBody))
	return sucRead, nil
}

func FakeUserJSON() ([]byte, error) {
	reqBody, err := json.Marshal(responses.BaseResponse{
		Success: true,
		Data:    FakeUser(),
		Error:   nil,
	})
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}
