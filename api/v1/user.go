package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Account" method:"post" summary:"Login"`
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
	g.Meta          `path:"/register" tags:"Account" method:"post" summary:"Register"`
	Username        string `v:"required#Please input username"`
	Password        string `v:"required#Please input password"`
	ConfirmPassword string `v:"required#Please input confirm password"`
	Email           string `v:"required#Please input email"`
	Gender          string `v:"required#Please input gender"`
	Phone           string `v:"required#Please input phone"`
}

type RegisterRes struct{}

type ProfileReq struct {
	g.Meta `path:"/profile" tags:"Account" method:"post" summary:"Get user's profile'"`
}

type ProfileRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   struct {
		Username string
		Gender   string
		Role     uint
		Email    string
		Phone    string
		Avatar   string
	}
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"Account" method:"get" summary:"Log out'"`
}

type LogoutRes struct {
}
