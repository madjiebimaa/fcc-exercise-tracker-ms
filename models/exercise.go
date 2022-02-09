package models

import (
	"context"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Description string             `json:"description" bson:"description"`
	Duration    int                `json:"duration" bson:"duration"`
	Date        time.Time          `json:"date" bson:"date"`
}

type ExerciseRepository interface {
	Create(ctx context.Context, exercise *Exercise) error
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]Exercise, error)
}

type ExerciseService interface {
	Create(ctx context.Context, req *requests.ExerciseCreate) (responses.ExerciseCreate, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) (responses.ExerciseLogs, error)
}
