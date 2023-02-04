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
	IOrder interface {
		CreateOrder(ctx context.Context, input model.CreateOrderForm) (response *model.ResponseOrderForm, err error)
		GenerateOrderCode() string
		ValidateTime(ctx context.Context, input model.CreateOrderForm) (res bool, err error)
		GetOrdersByUserId(ctx context.Context, userId int) (res []*entity.Order, err error)
		GetOrdersByPlaceId(ctx context.Context, placeId int) (res []*entity.Order, err error)
		GetOrderByOrderCode(ctx context.Context, orderCode string) (res *entity.Order, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
