package responses

type ExerciseCreate struct {
	User     BaseUser     `json:"user"`
	Exercise BaseExercise `json:"exercise"`
}

type ExerciseLogs struct {
	User      BaseUser       `json:"user"`
	Count     int            `json:"Count"`
	Exercises []BaseExercise `json:"exercises"`
}
