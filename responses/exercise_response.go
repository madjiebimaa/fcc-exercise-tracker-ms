package responses

type ExerciseCreate struct {
	User     BaseUser     `json:"user"`
	Exercise BaseExercise `json:"exercise"`
}

type ExerciseLogs struct {
	User      BaseUser       `json:"user"`
	Length    int            `json:"length"`
	Exercises []BaseExercise `json:"exercises"`
}
