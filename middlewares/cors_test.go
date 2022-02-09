package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/middlewares"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c, r := gin.CreateTestContext(rec)

	h := middlewares.CORS()
	r.ServeHTTP(rec, req)
	h(c)

	assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
}
