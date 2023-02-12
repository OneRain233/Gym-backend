package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Bank = cBank{}
var BankAdmin = cBankAdmin{}

type cBank struct{}
type cBankAdmin struct{}

func (c *cBankAdmin) AddBank(ctx context.Context, req *v1.AddBankReq) (res *v1.AddBankRes, err error) {
	res = &v1.AddBankRes{}
	form := model.AddBankForm{
		Name: req.Name,
	}
	err = service.Bank().AddBank(ctx, &form)

	if err != nil {
		return
	}
	return
}

func (c *cBankAdmin) UpdateBank(ctx context.Context, req *v1.UpdateBankReq) (res *v1.UpdateBankRes, err error) {
	res = &v1.UpdateBankRes{}
	form := model.UpdateBankForm{
		Id:   req.Id,
		Name: req.Name,
	}
	err = service.Bank().UpdateBank(ctx, &form)

	if err != nil {
		return
	}
	return
}

func (c *cBank) GetBanks(ctx context.Context, req *v1.GetBanksReq) (res *v1.GetBanksRes, err error) {
	res = &v1.GetBanksRes{}
	banks, err := service.Bank().GetBanks(ctx)

	if err != nil {
		return
	}
	res.Data = banks
	return
}

func (c *cBankAdmin) DeleteBank(ctx context.Context, req *v1.DeleteBankReq) (res *v1.DeleteBankRes, err error) {
	res = &v1.DeleteBankRes{}
	err = service.Bank().DeleteBank(ctx, req.Id)

	if err != nil {
		return
	}
	return
}
