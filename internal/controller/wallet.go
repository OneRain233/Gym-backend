package controller

import (
	v1 "Gym-backend/api/v1"
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
