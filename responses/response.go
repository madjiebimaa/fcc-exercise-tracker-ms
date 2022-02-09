package responses

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseExercise struct {
	ID          primitive.ObjectID `json:"id"`
	Description string             `json:"description"`
	Duration    int                `json:"duration"`
	Date        string             `json:"date"`
}

type BaseUser struct {
	ID       primitive.ObjectID `json:"id"`
	UserName string             `json:"username"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *FieldError `json:"error"`
}
