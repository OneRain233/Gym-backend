// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"context"
)

type (
	IUser interface {
		Login(ctx context.Context, input model.UserLoginForm) error
		Register(ctx context.Context, input model.UserRegisterForm) error
		CheckUserEmailFormat(email string) error
		CheckUserPhoneFormat(phone string) error
		GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.User, err error)
		GetUserByID(ctx context.Context, id uint) (user *entity.User, err error)
		ValidateUsername(ctx context.Context, username string) error
		ValidatePhone(ctx context.Context, phone string) error
		ValidateEmail(ctx context.Context, email string) error
		ValidateGender(ctx context.Context, gender int) error
		ValidateRole(ctx context.Context, role int) error
		UpdateEmptyAvatarPath(ctx context.Context, user *entity.User) error
		UpdateAvatar(ctx context.Context, userId uint, avatar string) error
		GetCurrentUser(ctx context.Context) (user *entity.User)
		UpdatePassword(ctx context.Context, user *entity.User, newPassword string, oldPassword string) error
		GetAllUser(ctx context.Context, pagination *model.Pagination) (users []*entity.User, err error)
		GetUserBySearch(ctx context.Context, search string) (users []*entity.User, err error)
		GetUserById(ctx context.Context, id uint) (user *entity.User, err error)
		UpdateUser(ctx context.Context, form *model.UserUpdateForm) error
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
