package fakes

import (
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	UserIDStr = "62034f817663d57f03323508"
	UserID, _ = primitive.ObjectIDFromHex(UserIDStr)
	UserName  = "madjiebimaa"

	ExerciseIDStr = "62034fe2fceb658638410a09"
	ExerciseID, _ = primitive.ObjectIDFromHex(ExerciseIDStr)
	Description   = "finished backend freecodecamp courses"
	Duration      = 3
	Date          = "2045-02-09"
	DateTime, _   = helpers.StripedDateToTime(Date)
	StrDate       = DateTime.Format(time.RFC1123)

	From  = "2044-02-09"
	To    = "2046-02-09"
	Limit = 1
)
