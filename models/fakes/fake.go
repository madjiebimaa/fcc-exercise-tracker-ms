package fakes

import (
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	user_id  = primitive.NewObjectID()
	username = "madjiebimaa"

	exercise_id = primitive.NewObjectID()
	description = "finished backend freecodecamp courses"
	duration    = 3
	date        = "2022-02-09"
	dateTime, _ = helpers.StripedDateToTime(date)
	strDate     = dateTime.Format(time.RFC1123)
)
