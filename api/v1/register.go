package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta          `path:"/register" tags:"Register" method:"post" summary:"Register"`
	Username        string `v:"required#Please input username"`
	Password        string `v:"required#Please input password"`
	ConfirmPassword string `v:"required#Please input confirm password"`
	Email           string `v:"required#Please input email"`
	Gender          string `v:"required#Please input gender"`
	Phone           string `v:"required#Please input phone"`
}

type RegisterRes struct {
	// json response
	g.Meta `mime:"application/json" example:"string"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Username string `json:"username"`
	}
}
