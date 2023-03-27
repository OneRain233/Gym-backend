// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	IOrder interface {
		CalculateAmount(ctx context.Context, input model.CreateOrderForm) (amount float64, discount float64, err error)
		CreateOrder(ctx context.Context, input model.CreateOrderForm) (response *model.ResponseOrderForm, err error)
		GenerateOrderCode() string
		ValidateTime(ctx context.Context, input model.CreateOrderForm) (res bool, err error)
		GetOrdersByUserId(ctx context.Context, userId int) (res []*entity.Order, err error)
		GetOrdersByPlaceId(ctx context.Context, placeId int) (res []*entity.Order, err error)
		GetOrderByOrderCode(ctx context.Context, orderCode string) (res *entity.Order, err error)
		GetAllOrders(ctx context.Context, pagination *model.Pagination) (res []*entity.Order, err error)
		GetRefundedOrder(ctx context.Context) (res []*entity.Order, err error)
		GetOrderByTimeRange(ctx context.Context, startTime *gtime.Time, endTime *gtime.Time) (res []*entity.Order, err error)
		GenerateOrderReceipt(ctx context.Context, orderCode string) (path string, err error)
		RefundOrder(ctx context.Context, orderCode string) (err error)
		GenerateQrSignature(qrContent map[string]interface{}) string
		CheckSignature(qrContent map[string]interface{}, sign string) bool
		StartOrder(ctx context.Context, orderCode string) (err error)
		EndOrder(ctx context.Context, orderCode string) (err error)
		CancelOrder(ctx context.Context, orderCode string) (err error)
		GetDailyOrderIncome(ctx context.Context, date *gtime.Time) (income float64, err error)
		CheckExpiredOrder(ctx context.Context) (err error)
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
