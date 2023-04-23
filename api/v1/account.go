package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/user/login" tags:"Account" method:"post" summary:"Login"`
	Username string `v:"required#Please input username"`
	Password string `v:"required#Please input password"`
}

type LoginRes struct {
	// json response
	g.Meta `mime:"application/json" example:"string"`
	Data   struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Role     int    `json:"role"`
		Token    string `json:"token"`
		IsCoach  bool   `json:"is_coach"`
	} `json:"data"`
}

type RegisterReq struct {
	g.Meta          `path:"/user/register" tags:"Account" method:"post" summary:"Register"`
	Username        string `v:"required#Please input username"`
	Password        string `v:"required#Please input password"`
	ConfirmPassword string `v:"required#Please input confirm password"`
	Email           string `v:"required#Please input email"`
	Gender          uint   `v:"required#Please input gender"`
	Phone           string `v:"required#Please input phone"`
}

type RegisterRes struct{}

type ProfileReq struct {
	g.Meta `path:"/user/profile" tags:"Account" method:"post" summary:"Get user's profile"`
}

type ProfileRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   struct {
		Username string `json:"username"`
		Gender   uint   `json:"gender"`
		Role     uint   `json:"role"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		IsCoach  bool   `json:"is_coach"`
	} `json:"data"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" tags:"Account" method:"get" summary:"Log out"`
}

type LogoutRes struct {
}

type ChangePasswdReq struct {
	g.Meta          `path:"/change-passwd" tags:"Account" method:"post" summary:"Change password"`
	OldPassword     string `v:"required#Please input old password" json:"old_password"`
	NewPassword     string `v:"required#Please input new password" json:"new_password"`
	ConfirmPassword string `v:"required#Please input confirm password" json:"confirm_password"`
}

type ChangePasswdRes struct {
}

type GetUserListReq struct {
	g.Meta `path:"/user/list" tags:"Account" method:"post" summary:"Get user list"`
	Limit  int `json:"limit"`
	Page   int `json:"page"`
}

type GetUserListRes struct {
	g.Meta `mime:"application/json" example:"string"`
	User   []*entity.User `json:"user"`
}

type GetUserSearchReq struct {
	g.Meta   `path:"/user/search" tags:"Account" method:"post" summary:"Get user by username"`
	Username string `v:"required#Please input username"`
}

type GetUserSearchRes struct {
	g.Meta `mime:"application/json" example:"string"`
	User   []*entity.User `json:"user"`
}

type GetUserByIdReq struct {
	g.Meta `path:"/user/get" tags:"Account" method:"post" summary:"Get user by id"`
	Id     uint
}

type GetUserByIdRes struct {
	g.Meta `mime:"application/json" example:"string"`
	User   *entity.User `json:"user"`
}

type UpdateUserReq struct {
	g.Meta   `path:"/user/update" tags:"Account" method:"post" summary:"Update user"`
	Id       uint   `json:"id"`
	Role     uint   `json:"role"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   uint   `json:"gender"`
	IsCoach  bool   `json:"is_coach"`
}

type UserUpdateUserReq struct {
	g.Meta `path:"/user/profile/update" tags:"Account" method:"post" summary:"Normal User Update user"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Gender uint   `json:"gender"`
}

type UserUpdateUserRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type UpdateUserRes struct {
}

type UpdateAvatarReq struct {
	g.Meta `path:"/user/avatar/update" tags:"Account" method:"post" summary:"Update user's avatar"`
	Avatar string `v:"required#Please input avatar"`
}

type UpdateAvatarRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type ForgetPasswordReq struct {
	g.Meta `path:"/user/forget-password" tags:"Account" method:"post" summary:"Forget password"`
	Email  string `json:"email" v:"required#Please input email"`
}

type ForgetPasswordRes struct {
	// TODO: send email instead of return token
	Token string `json:"token"`
}

type ResetPasswordReq struct {
	g.Meta          `path:"/user/reset-password" tags:"Account" method:"post" summary:"Reset password"`
	Password        string `json:"password" v:"required#Please input password"`
	ConfirmPassword string `json:"confirm_password" v:"required#Please input confirm password"`
	Token           string `json:"token" v:"required#Please input token"`
}

type ResetPasswordRes struct{}

type ValidateTokenReq struct {
	g.Meta `path:"/user/validate-token" tags:"Account" method:"post" summary:"Validate token"`
	Token  string `json:"token" v:"required#Please input token"`
}

type ValidateTokenRes struct {
	Validate bool `json:"validate"`
}
