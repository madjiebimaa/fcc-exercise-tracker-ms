package requests

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseCreate struct {
	UserID      primitive.ObjectID `json:"user_id"`
	Description string             `json:"description" binding:"required"`
	Duration    int                `json:"duration" binding:"required,gte=1"`
	Date        string             `json:"date" binding:"required"`
}
