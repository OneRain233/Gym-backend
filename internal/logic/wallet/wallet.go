package wallet

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sWallet struct {
}

func New() *sWallet {
	return &sWallet{}
}

func init() {
	service.RegisterWallet(New())
}

func (s *sWallet) GetWallet(ctx context.Context) (wallet *entity.Wallet, err error) {
	userId := service.Session().GetUser(ctx).Id
	err = dao.Wallet.Ctx(ctx).Where("userId", userId).Scan(&wallet)
	if err != nil {
		return
	}
	return
}

func (s *sWallet) SetStatus(ctx context.Context, status int) error {
	userId := service.Session().GetUser(ctx).Id
	var wallet entity.Wallet
	err := dao.Wallet.Ctx(ctx).Where("userId", userId).Scan(&wallet)
	if err != nil {
		return err
	}
	wallet.Status = status
	_, err = dao.Wallet.Ctx(ctx).Where("userId", userId).Update(&wallet)
	return err
}