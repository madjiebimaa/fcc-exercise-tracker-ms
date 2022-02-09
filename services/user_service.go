package services

import (
	"context"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
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

func (u *userService) Register(c context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.Register(ctx, user)
}
