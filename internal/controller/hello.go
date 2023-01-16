package controller

import (
	"Gym-backend/internal/service"
	"context"

	v1 "Gym-backend/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	user := service.Session().GetUser(ctx)
	username := user.Username
	res = &v1.HelloRes{}
	res.Msg = "hello world " + username

	return
}
