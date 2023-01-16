package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Register = cRegister{}

type cRegister struct{}

func (c *cRegister) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	res = &v1.RegisterRes{}
	if err := service.User().Register(ctx, model.UserRegisterForm{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Gender:   req.Gender,
	}); err != nil {
		res.Status = 500
		res.Msg = err.Error()
	}
	res.Status = 200
	res.Msg = "Register success"
	res.Data.Username = req.Username
	return

}
