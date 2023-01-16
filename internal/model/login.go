package model

type UserLoginForm struct {
	Username string
	Password string
}
type UserRegisterForm struct {
	Username        string
	Password        string
	ConfirmPassword string
	Email           string
	Gender          string
	Phone           string
}
