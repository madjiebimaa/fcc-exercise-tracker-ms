package fakes

import (
	"encoding/json"
	"strings"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
)

func FakeUserRegisterRequest() requests.UserRegister {
	return requests.UserRegister{
		UserName: username,
	}
}

func FakeUser() models.User {
	return models.User{
		ID:       user_id,
		UserName: username,
	}
}

func FakeUserJSON() ([]byte, error) {
	reqBody, err := json.Marshal(FakeUser())
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}

func FakeUserReader(username string) (*strings.Reader, error) {
	reqBody, err := json.Marshal(models.User{
		UserName: username,
	})
	if err != nil {
		return nil, err
	}
	sucRead := strings.NewReader(string(reqBody))
	return sucRead, nil
}
