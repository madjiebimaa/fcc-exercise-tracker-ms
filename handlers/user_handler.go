package handlers

import (
	"net/http"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user input is not valid",
		})
		return
	}

	user := models.User{
		ID:       primitive.NewObjectID(),
		UserName: req.UserName,
	}

	ctx := c.Request.Context()
	if err := u.userService.Register(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
