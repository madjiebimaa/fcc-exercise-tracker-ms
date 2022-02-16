package repositories

import (
	"context"
	"log"
	"time"

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

func filterDate(ctx context.Context) ([]primitive.E, *options.FindOptions) {
	var filterRequest requests.ExerciseLogsFilter
	helpers.QueryToStruct(ctx, constants.LOGS_FILTER, &filterRequest)

	filter := []bson.E{}
	var from time.Time
	if filterRequest.From != "" {
		from, _ = helpers.StripedDateToTime(filterRequest.From)
	} else {
		from = time.Now()
	}
	filterFrom := bson.E{
		Key:   "$gte",
		Value: primitive.NewDateTimeFromTime(from),
	}
	filter = append(filter, filterFrom)

	var to time.Time
	if filterRequest.To != "" {
		to, _ = helpers.StripedDateToTime(filterRequest.To)
	} else {
		to = from.Add(time.Hour * 24 * 60 * 12 * 10)
	}
	filterTo := bson.E{
		Key:   "$lt",
		Value: primitive.NewDateTimeFromTime(to),
	}
	filter = append(filter, filterTo)

	options := options.Find()
	if filterRequest.Limit != 0 {
		options.SetLimit(int64(filterRequest.Limit))
	}

	return filter, options
}

func (m *mongoExerciseRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Exercise, error) {
	filterLogs, ops := filterDate(ctx)
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "date", Value: bson.D{
			filterLogs[0],
			filterLogs[1],
		}},
	}

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
