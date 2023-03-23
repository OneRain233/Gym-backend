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
	ICard interface {
		GetCard(ctx context.Context) (card *entity.WalletCard, err error)
		ValidateCard(ctx context.Context, input *model.BindCardForm) error
		ValidatePhone(ctx context.Context, input *model.BindCardForm) error
		BindCard(ctx context.Context, input *model.BindCardForm) error
		GetCardsByUserId(ctx context.Context, userId int) (cards []*entity.WalletCard, err error)
		GetCardsCountByUserId(ctx context.Context, userId int) (count int, err error)
		Pay(ctx context.Context, input *model.CardPayForm) error
		PayForSubscription(ctx context.Context, input *model.CardPayForm) error
		Recharge(ctx context.Context, input *model.CardRechargeForm) error
		DeleteCard(ctx context.Context, cardId int) error
	}
)

var (
	localCard ICard
)

func Card() ICard {
	if localCard == nil {
		panic("implement not found for interface ICard, forgot register?")
	}
	return localCard
}

func RegisterCard(i ICard) {
	localCard = i
}
