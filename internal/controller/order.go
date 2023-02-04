package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Order = cOrder{}
var OrderAdmin = cOrderAdmin{}

type cOrder struct{}
type cOrderAdmin struct{}

func (c *cOrder) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error) {
	res = &v1.CreateOrderRes{}
	form := model.CreateOrderForm{
		UserId:    service.Session().GetUser(ctx).Id,
		PlaceId:   req.PlaceId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	resp, err := service.Order().CreateOrder(ctx, form)
	if err != nil {
		return
	}

	res.OrderCode = resp.OrderCode
	res.Amount = resp.Amount
	res.PlaceId = resp.PlaceId
	res.StartTime = resp.StartTime
	res.EndTime = resp.EndTime

	return
}
