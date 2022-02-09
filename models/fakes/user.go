package fakes

import (
	"encoding/json"
	"strings"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
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

func FakeUserReader(UserName string) (*strings.Reader, error) {
	reqBody, err := json.Marshal(models.User{
		UserName: UserName,
	})
	if err != nil {
		return nil, err
	}
	sucRead := strings.NewReader(string(reqBody))
	return sucRead, nil
}

func FakeUserJSON() ([]byte, error) {
	reqBody, err := json.Marshal(FakeUser())
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}
