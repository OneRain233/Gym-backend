package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Wallet = cWallet{}

type cWallet struct {
}

func (c *cWallet) GetWalletInfo(ctx context.Context, req *v1.GetWalletReq) (res *v1.GetWalletRes, err error) {
	res = &v1.GetWalletRes{}
	res.Info, err = service.Wallet().GetFullWalletInfo(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cWallet) TopUp(ctx context.Context, req *v1.TopUpReq) (res *v1.TopUpRes, err error) {
	res = &v1.TopUpRes{}
	input := &model.WalletTopUpForm{
		Amount: req.Amount,
		CardId: req.CardId,
	}
	err = service.Wallet().TopUp(ctx, input)
	if err != nil {
		return
	}
	return
}

func (c *cWallet) WithDraw(ctx context.Context, req *v1.WithDrawReq) (res *v1.WithDrawRes, err error) {
	res = &v1.WithDrawRes{}
	input := &model.WalletWithdrawForm{
		Amount: req.Amount,
		CardId: req.CardId,
	}
	err = service.Wallet().WithDraw(ctx, input)
	if err != nil {
		return
	}
	return
}
