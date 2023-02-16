// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"context"
)

type (
	IWallet interface {
		GetWallet(ctx context.Context) (wallet *entity.Wallet, err error)
		Pay(ctx context.Context, input *model.WalletPayForm) error
		GetFullWalletInfo(ctx context.Context) (walletInfo *model.WalletInfo, err error)
		CreateWallet(ctx context.Context) error
		CreateWalletForUser(ctx context.Context, userId int) error
		SetStatus(ctx context.Context, status int) error
		GetCardsInWallet(ctx context.Context) (cards []*entity.WalletCard, err error)
		Refund(ctx context.Context, order *entity.Order) error
	}
)

var (
	localWallet IWallet
)

func Wallet() IWallet {
	if localWallet == nil {
		panic("implement not found for interface IWallet, forgot register?")
	}
	return localWallet
}

func RegisterWallet(i IWallet) {
	localWallet = i
}
