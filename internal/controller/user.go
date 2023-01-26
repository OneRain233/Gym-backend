package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var User = cUser{}

type cUser struct {
}

func (c *cUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	res = &v1.RegisterRes{}
	err = service.User().Register(ctx, model.UserRegisterForm{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Gender:   req.Gender,
	})
	return

}

func (c *cUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	err = service.User().Login(ctx, model.UserLoginForm{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	userEntity := service.Session().GetUser(ctx)
	res.Data.Username = userEntity.Username
	res.Data.Avatar = userEntity.Avatar
	res.Data.Role = userEntity.Role
	return
}

func (c *cUser) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res = &v1.LogoutRes{}
	err = service.Session().RemoveUser(ctx)
	if err != nil {
		return
	}

	return
}
