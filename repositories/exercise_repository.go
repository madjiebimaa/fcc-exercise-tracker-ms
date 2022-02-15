package repositories

import (
	"context"
	"log"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/constants"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/helpers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		log.Println(err)
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

func filterDate(ctx context.Context) ([]primitive.E, *options.FindOptions) {
	var filterRequest requests.ExerciseLogsFilter
	helpers.QueryToStruct(ctx, constants.LOGS_FILTER, &filterRequest)

	filter := []bson.E{}
	if filterRequest.From != "" {
		from, _ := helpers.StripedDateToTime(filterRequest.From)
		filterFrom := bson.E{
			Key:   "$gte",
			Value: primitive.NewDateTimeFromTime(from),
		}
		filter = append(filter, filterFrom)
	}

	if filterRequest.To != "" {
		to, _ := helpers.StripedDateToTime(filterRequest.To)
		filterTo := bson.E{
			Key:   "$lt",
			Value: primitive.NewDateTimeFromTime(to),
		}
		filter = append(filter, filterTo)
	}

	options := options.Find()
	if filterRequest.Limit != 0 {
		options.SetLimit(int64(filterRequest.Limit))
	}

	return filter, options
}

// TODO: Implement filtering base on query param From, To, and Limit
func (m *mongoExerciseRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Exercise, error) {
	// var filterRequest requests.ExerciseLogsFilter
	// helpers.QueryToStruct(ctx, constants.LOGS_FILTER, &filterRequest)
	// from, _ := helpers.StripedDateToTime(filterRequest.From)
	// to, _ := helpers.StripedDateToTime(filterRequest.To)
	filterLogs, ops := filterDate(ctx)
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "date", Value: bson.D{
			filterLogs[0],
			filterLogs[1],
			// {Key: "$gte", Value: primitive.NewDateTimeFromTime(from)},
			// {Key: "$lt", Value: primitive.NewDateTimeFromTime(to)},
		}},
	}

	// options := options.Find()
	// options.SetLimit(int64(filterRequest.Limit))

	cur, err := m.coll.Find(ctx, filter, ops)
	if err == nil {
		var exercises []models.Exercise
		for cur.Next(ctx) {
			var exercise models.Exercise
			if err := cur.Decode(&exercise); err != nil {
				log.Println(err)
				return []models.Exercise{}, models.ErrInternalServerError
			}
			exercises = append(exercises, exercise)
		}

		return exercises, nil
	} else if err == mongo.ErrNoDocuments {
		return []models.Exercise{}, nil
	} else {
		log.Println(err)
		return []models.Exercise{}, models.ErrInternalServerError
	}
}
