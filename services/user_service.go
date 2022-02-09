package services

import (
	"context"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	userRepo       models.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepo models.UserRepository, contextTimeout time.Duration) models.UserService {
	return &userService{
		userRepo,
		contextTimeout,
	}
}

func (u *userService) Register(c context.Context, req *requests.UserRegister) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user := models.User{
		ID:       primitive.NewObjectID(),
		UserName: req.UserName,
	}

	if err := u.userRepo.Register(ctx, &user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
