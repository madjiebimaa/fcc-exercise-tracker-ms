package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/stretchr/testify/assert"
)

func CreateUserRequest(username string) (*strings.Reader, error) {
	reqBody, err := json.Marshal(models.User{
		UserName: username,
	})
	if err != nil {
		return nil, err
	}
	sucRead := strings.NewReader(string(reqBody))
	return sucRead, nil
}

func TestUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	sucReq, err := CreateUserRequest("madjiebimaa")
	assert.NoError(t, err)

	failReq, err := CreateUserRequest("adjie")
	assert.NoError(t, err)

	sucRes, err := json.Marshal(models.User{
		UserName: "madjiebimaa",
	})
	assert.NoError(t, err)

	failRes, err := json.Marshal(gin.H{
		"message": "user input is not valid",
	})

	testCases := []struct {
		name       string
		reqBody    io.Reader
		statusCode int
		resBody    []byte
	}{
		{
			name:       "success",
			reqBody:    sucReq,
			statusCode: http.StatusCreated,
			resBody:    sucRes,
		},
		{
			name:       "fail user input is not valid",
			reqBody:    failReq,
			statusCode: http.StatusBadRequest,
			resBody:    failRes,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/users", tt.reqBody)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			_, r := gin.CreateTestContext(rec)
			// handlers.NewUserHandler()
			// r.POST("/api/users", handlers.UserHandler)

			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
			assert.Equal(t, tt.resBody, rec.Body.Bytes())
		})
	}
}
