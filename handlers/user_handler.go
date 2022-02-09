package handlers

import (
	"net/http"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/helpers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService models.UserService
}

func NewUserHandler(userService models.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (u *UserHandler) Register(c *gin.Context) {
	var req requests.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "request body", models.ErrBadInput)
		return
	}

	ctx := c.Request.Context()
	user, err := u.userService.Register(ctx, &req)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
		return
	}

	helpers.SuccessResponse(c, http.StatusCreated, user)
}
