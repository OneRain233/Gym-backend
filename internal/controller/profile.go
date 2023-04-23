package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/service"
	"context"
)

var Profile = cProfile{}

type cProfile struct{}

func (c *cProfile) GetProfile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res = &v1.ProfileRes{}
	user := service.Session().GetUser(ctx)
	res.Data.Username = user.Username
	res.Data.Avatar = user.Avatar
	res.Data.Role = uint(user.Role)
	res.Data.Email = user.Email
	res.Data.Phone = user.Phone
	res.Data.Gender = uint(user.Gender)
	if user.IsCoach == 1 {
		res.Data.IsCoach = true
	} else {
		res.Data.IsCoach = false
	}
	return

}
