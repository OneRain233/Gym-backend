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
	}
}

type RegisterReq struct {
	g.Meta          `path:"/user/register" tags:"Account" method:"post" summary:"Register"`
	Username        string `v:"required#Please input username"`
	Password        string `v:"required#Please input password"`
	ConfirmPassword string `v:"required#Please input confirm password"`
	Email           string `v:"required#Please input email"`
	Gender          string `v:"required#Please input gender"`
	Phone           string `v:"required#Please input phone"`
}

type RegisterRes struct{}

type ProfileReq struct {
	g.Meta `path:"/user/profile" tags:"Account" method:"post" summary:"Get user's profile"`
}

type ProfileRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   struct {
		Username string
		Gender   uint
		Role     uint
		Email    string
		Phone    string
		Avatar   string
	}
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" tags:"Account" method:"get" summary:"Log out"`
}

type LogoutRes struct {
}

type ChangePasswdReq struct {
	g.Meta          `path:"/change-passwd" tags:"Account" method:"post" summary:"Change password"`
	OldPassword     string `v:"required#Please input old password"`
	NewPassword     string `v:"required#Please input new password"`
	ConfirmPassword string `v:"required#Please input confirm password"`
}

type ChangePasswdRes struct {
}

type GetUserListReq struct {
	g.Meta `path:"/user/list" tags:"Account" method:"post" summary:"Get user list"`
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
