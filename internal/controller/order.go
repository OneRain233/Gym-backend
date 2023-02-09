package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
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

func (c *cOrderAdmin) GetAllOrder(ctx context.Context, req *v1.GetAllOrderReq) (res *v1.GetAllOrderRes, err error) {
	res = &v1.GetAllOrderRes{}
	res.Order, err = service.Order().GetAllOrders(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cOrderAdmin) GetOrderById(ctx context.Context, req *v1.GetOrderByUserIdReq) (res *v1.GetOrderByUserIdRes, err error) {
	res = &v1.GetOrderByUserIdRes{}
	res.Order, err = service.Order().GetOrdersByUserId(ctx, req.UserId)
	if err != nil {
		return
	}
	if res.Order == nil {
		err = gerror.New("order not found")
	}
	return
}

func (c *cOrderAdmin) GetOrderByPlaceId(ctx context.Context, req *v1.GetOrderByPlaceIdReq) (res *v1.GetOrderByPlaceIdRes, err error) {
	res = &v1.GetOrderByPlaceIdRes{}
	res.Order, err = service.Order().GetOrdersByPlaceId(ctx, req.PlaceId)
	if err != nil {
		return
	}
	if res.Order == nil {
		err = gerror.New("order not found")
	}
	return
}

func (c *cOrderAdmin) GetOrderByOrderCode(ctx context.Context, req *v1.GetOrderByOrderCodeReq) (res *v1.GetOrderByOrderCodeRes, err error) {
	res = &v1.GetOrderByOrderCodeRes{}
	res.Order, err = service.Order().GetOrderByOrderCode(ctx, req.OrderCode)
	if err != nil {
		return
	}
	if res.Order == nil {
		err = gerror.New("order not found")
	}
	return
}

func (c *cOrderAdmin) GetOrderByTime(ctx context.Context, req *v1.GetOrderByTimeReq) (res *v1.GetOrderByTimeRes, err error) {
	res = &v1.GetOrderByTimeRes{}
	startTime := gtime.NewFromStr(req.StartTime)
	endTime := gtime.NewFromStr(req.EndTime)
	res.Order, err = service.Order().GetOrderByTimeRange(ctx, startTime, endTime)
	if err != nil {
		return
	}
	if res.Order == nil {
		err = gerror.New("order not found")
	}

	return
}

func (c *cOrder) GetReceipt(ctx context.Context, req *v1.GetReceiptReq) (res *v1.GetReceiptRes, err error) {
	res = &v1.GetReceiptRes{}
	path, err := service.Order().GenerateOrderReceipt(ctx, req.OrderCode)
	if err != nil {
		return
	}
	// return the file as a response
	g.RequestFromCtx(ctx).Response.ServeFile(path)
	return
}
