package models

import (
	"context"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	UserName string             `json:"username" bson:"username"`
}

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	GetByID(ctx context.Context, userID primitive.ObjectID) (User, error)
}

type UserService interface {
	Register(ctx context.Context, req *requests.UserRegister) (User, error)
}
