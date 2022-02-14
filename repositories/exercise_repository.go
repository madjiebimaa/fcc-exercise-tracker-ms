package repositories

import (
	"context"
	"log"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// func (r *mongoExerciseRepository) queryRowsFilter(ctx context.Context, tx *sql.Tx, sql string) (*sql.Rows, error) {
// 	var filterValues []interface{}
// 	var filterRequest requests.ExerciseCreateFilter
// 	helpers.QueryToStruct(ctx, constants.LOGS_FILTER, &filterRequest)

// 	if filterRequest.MinLength != 0 {
// 		filterValues = append(filterValues, filterRequest.MinLength)
// 		sql += ` HAVING content_length > ?`
// 	}

// 	if filterRequest.MaxLength != 0 {
// 		filterValues = append(filterValues, filterRequest.MaxLength)
// 		sql += ` AND content_length < ?`
// 	}

// 	if filterRequest.Limit != 0 {
// 		filterValues = append(filterValues, filterRequest.Limit)
// 		sql += ` LIMIT ?`
// 	}

// 	if filterRequest.Offset != 0 {
// 		filterValues = append(filterValues, filterRequest.Offset)
// 		sql += ` OFFSET ?`
// 	}

// 	return tx.QueryContext(ctx, sql, filterValues...)
// }

func (m *mongoExerciseRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Exercise, error) {
	filter := bson.D{{Key: "user_id", Value: userID}}
	cur, err := m.coll.Find(ctx, filter)
	if err == nil {
		var exercises []models.Exercise
		for cur.Next(ctx) {
			var exercise models.Exercise
			if err := cur.Decode(&exercise); err != nil {
				log.Fatal(err)
				return []models.Exercise{}, models.ErrInternalServerError
			}
			exercises = append(exercises, exercise)
		}

		return exercises, nil
	} else if err == mongo.ErrNoDocuments {
		return []models.Exercise{}, nil
	} else {
		log.Fatal(err)
		return []models.Exercise{}, models.ErrInternalServerError
	}
}
