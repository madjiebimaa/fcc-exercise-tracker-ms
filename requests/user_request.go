package requests

type UserRegister struct {
	UserName string `json:"username" binding:"required,min=5"`
}
