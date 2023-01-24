package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"Login"`
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
	g.Meta          `path:"/register" tags:"Register" method:"post" summary:"Register"`
	Username        string `v:"required#Please input username"`
	Password        string `v:"required#Please input password"`
	ConfirmPassword string `v:"required#Please input confirm password"`
	Email           string `v:"required#Please input email"`
	Gender          string `v:"required#Please input gender"`
	Phone           string `v:"required#Please input phone"`
}

type RegisterRes struct{}
