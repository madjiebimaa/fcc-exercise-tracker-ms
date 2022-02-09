package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/helpers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseHandler struct {
	exerciseService models.ExerciseService
}

func NewExerciseHandler(exerciseService models.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{
		exerciseService,
	}
}

func (e *ExerciseHandler) Create(c *gin.Context) {
	u := c.Param("userID")
	if u == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "url param", models.ErrNotFound)
		return
	}

	var req requests.ExerciseCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "request body", models.ErrBadInput)
		return
	}

	userID, err := primitive.ObjectIDFromHex(u)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "url param", err)
		return
	}
	req.UserID = userID

	ctx := c.Request.Context()
	res, err := e.exerciseService.Create(ctx, &req)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
		return
	}

	helpers.SuccessResponse(c, http.StatusCreated, res)
}
