package responses

type ExerciseCreate struct {
	User     BaseUser     `json:"user"`
	Exercise BaseExercise `json:"exercise"`
}
