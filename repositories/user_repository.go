package repositories

import (
	"context"
	"log"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoUserRepository struct {
	coll *mongo.Collection
}

func NewMongoUserRepository(coll *mongo.Collection) models.UserRepository {
	return &mongoUserRepository{
		coll,
	}
}

func (m *mongoUserRepository) Register(ctx context.Context, user *models.User) error {
	m.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if _, err := m.coll.InsertOne(ctx, user); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
