package repositories

import (
	"context"
	"log"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoExerciseRepository struct {
	coll *mongo.Collection
}

func NewMongoExerciseRepository(coll *mongo.Collection) models.ExerciseRepository {
	return &mongoExerciseRepository{
		coll,
	}
}

func (m *mongoExerciseRepository) Create(ctx context.Context, exercise *models.Exercise) error {
	if _, err := m.coll.InsertOne(ctx, exercise); err != nil {
		log.Fatal(err)
		return models.ErrInternalServerError
	}

	return nil
}
