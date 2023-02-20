package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/service"
	"context"
)

var Credit = cCredit{}
var CreditAdmin = cCreditAdmin{}

type cCredit struct {
}
type cCreditAdmin struct{}

func (c *cCredit) GetCredit(ctx context.Context, req *v1.GetCreditReq) (res *v1.GetCreditRes, err error) {
	res = &v1.GetCreditRes{}
	res.Credit, err = service.Credit().GetCredit(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cCreditAdmin) GetAllCredit(ctx context.Context, req *v1.AdminGetCreditReq) (res *v1.AdminGetCreditRes, err error) {
	res = &v1.AdminGetCreditRes{}
	res.Credit, err = service.Credit().GetAllCredit(ctx)
	if err != nil {
		return
	}
	return
}
