package repositories

import (
	"context"
	"log"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	_, err := m.coll.InsertOne(ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		return models.ErrConflict
	} else if err != nil {
		log.Fatal(err)
		return models.ErrInternalServerError
	}

	return nil
}

func (m *mongoUserRepository) GetByID(ctx context.Context, userID primitive.ObjectID) (models.User, error) {
	filter := bson.D{{Key: "_id", Value: userID}}
	var user models.User
	if err := m.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		log.Fatal(err)
		return models.User{}, models.ErrInternalServerError
	}

	return user, nil
}
