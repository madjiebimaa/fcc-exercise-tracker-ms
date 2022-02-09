package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	UserName string             `json:"username" bson:"username"`
}

type UserRepository interface {
	Register(ctx context.Context, user *User) error
}

type UserService interface {
	Register(ctx context.Context, user *User) error
}
