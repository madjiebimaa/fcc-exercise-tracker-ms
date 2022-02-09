package services

import (
	"context"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/helpers"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/requests"
	"github.com/madjiebimaa/fcc-exercise-tracker-ms/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type exerciseService struct {
	exerciseRepo   models.ExerciseRepository
	userRepo       models.UserRepository
	contextTimeout time.Duration
}

func NewExerciseService(exerciseRepo models.ExerciseRepository, userRepo models.UserRepository, contextTimeout time.Duration) models.ExerciseService {
	return &exerciseService{
		exerciseRepo,
		userRepo,
		contextTimeout,
	}
}

func (e *exerciseService) Create(c context.Context, req *requests.ExerciseCreate) (responses.ExerciseCreate, error) {
	ctx, cancel := context.WithTimeout(c, e.contextTimeout)
	defer cancel()

	user, err := e.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return responses.ExerciseCreate{}, err
	}

	if user.ID == primitive.NilObjectID {
		return responses.ExerciseCreate{}, models.ErrNotFound
	}

	date, err := helpers.StripedDateToTime(req.Date)
	if err != nil {
		return responses.ExerciseCreate{}, err
	}

	exercise := models.Exercise{
		ID:          primitive.NewObjectID(),
		Description: req.Description,
		Duration:    req.Duration,
		Date:        date,
	}

	if err := e.exerciseRepo.Create(ctx, &exercise); err != nil {
		return responses.ExerciseCreate{}, err
	}

	res := responses.ExerciseCreate{
		User:     responses.BaseUser(user),
		Exercise: helpers.ToBaseExercise(exercise),
	}

	return res, nil
}

func (e *exerciseService) GetByUserID(c context.Context, userID primitive.ObjectID) (responses.ExerciseLogs, error) {
	ctx, cancel := context.WithTimeout(c, e.contextTimeout)
	defer cancel()

	user, err := e.userRepo.GetByID(ctx, userID)
	if err != nil {
		return responses.ExerciseLogs{}, err
	}

	if user.ID == primitive.NilObjectID {
		return responses.ExerciseLogs{}, models.ErrNotFound
	}

	exercises, err := e.exerciseRepo.GetByUserID(ctx, userID)
	if err != nil {
		return responses.ExerciseLogs{}, err
	}

	if len(exercises) == 0 {
		return responses.ExerciseLogs{}, models.ErrNotFound
	}

	res := responses.ExerciseLogs{
		User:      responses.BaseUser(user),
		Length:    len(exercises),
		Exercises: helpers.ToBaseExercises(exercises),
	}

	return res, nil
}
