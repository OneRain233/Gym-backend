package model

import "Gym-backend/internal/model/entity"

type UserLoginForm struct {
	Username string
	Password string
}
type UserRegisterForm struct {
	Username        string
	Password        string
	ConfirmPassword string
	Email           string
	Gender          uint
	Phone           string
}

type UserUpdateForm struct {
	Id       uint
	Username string
	Role     uint
	Email    string
	Phone    string
	Gender   uint
	IsCoach  bool
}

type UserProfileWithCredit struct {
	User   *entity.User
	Credit *entity.Credit
}
