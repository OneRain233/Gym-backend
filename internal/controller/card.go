package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Card = cCard{}

type cCard struct {
}

func (c *cCard) BindCard(ctx context.Context, req *v1.BindCardReq) (res *v1.BindCardRes, err error) {
	res = &v1.BindCardRes{}
	form := model.BindCardForm{
		BankId:      req.BankId,
		CardAccount: req.CardAccount,
		Phone:       req.Phone,
		Amount:      0.0,
	}
	err = service.Card().BindCard(ctx, &form)
	if err != nil {
		return
	}
	return
}
