package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/consts"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

var Order = cOrder{}
var OrderAdmin = cOrderAdmin{}
var OrderManager = cOrderManager{}

type cOrder struct{}
type cOrderAdmin struct{}
type cOrderManager struct{}

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

func (c *cOrder) CreateRegularOrder(ctx context.Context, req *v1.CreateRegularOrderReq) (res *v1.CreateRegularOrderRes, err error) {
	res = &v1.CreateRegularOrderRes{}
	form := model.CreateRegularOrderFormWeekly{
		UserId:           service.Session().GetUser(ctx).Id,
		PlaceId:          req.PlaceId,
		SessionStartTime: req.SessionStartTime,
		SessionEndTime:   req.SessionEndTime,
		StartDay:         req.StartDay,
		WeekCount:        req.WeekCount,
	}

	resp, err := service.Order().CreateRegularWeeklyOrder(ctx, form)
	if err != nil {
		return
	}
	res.OrderCode = resp.OrderCode
	res.Orders = resp
	return
}

func (c *cOrderManager) GetAllOrder(ctx context.Context, req *v1.GetAllOrderReq) (res *v1.GetAllOrderRes, err error) {
	res = &v1.GetAllOrderRes{}
	var orders []*entity.Order

	pagination := &model.Pagination{
		Limit: req.Limit,
		Page:  req.Page,
	}

	orders, err = service.Order().GetAllOrders(ctx, pagination)
	if err != nil {
		return
	}
	for _, order := range orders {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Order = append(res.Order, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
	}
	if err != nil {
		return
	}

	return
}

func (c *cOrderManager) GetOrderById(ctx context.Context, req *v1.GetOrderByUserIdReq) (res *v1.GetOrderByUserIdRes, err error) {
	res = &v1.GetOrderByUserIdRes{}
	var order []*entity.Order
	order, err = service.Order().GetOrdersByUserId(ctx, req.UserId)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}

	for _, order := range order {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Order = append(res.Order, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
	}

	return
}

func (c *cOrderManager) GetOrderByPlaceId(ctx context.Context, req *v1.GetOrderByPlaceIdReq) (res *v1.GetOrderByPlaceIdRes, err error) {
	res = &v1.GetOrderByPlaceIdRes{}
	var order []*entity.Order
	order, err = service.Order().GetOrdersByPlaceId(ctx, req.PlaceId)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}
	for _, order := range order {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Order = append(res.Order, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
	}

	return
}

func (c *cOrder) GetOrderByOrderCode(ctx context.Context, req *v1.GetOrderByOrderCodeReq) (res *v1.GetOrderByOrderCodeRes, err error) {
	res = &v1.GetOrderByOrderCodeRes{}
	var order *entity.Order
	order, err = service.Order().GetOrderByOrderCode(ctx, req.OrderCode)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}
	if order.UserId != service.Session().GetUser(ctx).Id {
		err = gerror.New("permission denied")
		return
	}
	place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
	if err1 != nil {
		err = err1
	}
	res.Order = &model.AdminResponseOrderForm{
		Order: order,
		Place: place,
	}

	return
}

func (c *cOrder) GetRegularOrderByParentOrderCode(ctx context.Context, req *v1.GetRegularOrderInfoReq) (res *v1.GetRegularOrderInfoRes, err error) {
	res = &v1.GetRegularOrderInfoRes{}
	var orders []*entity.Order
	orders, err = service.Order().GetRegularOrdersByOrderCode(ctx, req.OrderCode)
	if err != nil {
		return
	}
	if orders == nil {
		err = gerror.New("order not found")
		return
	}
	for _, order := range orders {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Orders = append(res.Orders, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
		if order.Status == consts.OrderStatusRefunded || order.Status == consts.OrderStatusCancelled {
			continue
		}
		res.Amount += order.Amount
	}
	res.OrderCode = req.OrderCode
	return
}

func (c *cOrderManager) GetOrderByTime(ctx context.Context, req *v1.GetOrderByTimeReq) (res *v1.GetOrderByTimeRes, err error) {
	res = &v1.GetOrderByTimeRes{}
	startTime := gtime.NewFromStr(req.StartTime)
	endTime := gtime.NewFromStr(req.EndTime)
	var orders []*entity.Order
	orders, err = service.Order().GetOrderByTimeRange(ctx, startTime, endTime)
	if orders == nil {
		err = gerror.New("order not found")
		return
	}

	if err != nil {
		return
	}
	for _, order := range orders {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Order = append(res.Order, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
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

func (c *cOrder) Refund(ctx context.Context, req *v1.RefundOrderReq) (res *v1.RefundOrderRes, err error) {
	res = &v1.RefundOrderRes{}
	err = service.Order().RefundOrder(ctx, req.OrderCode)
	if err != nil {
		return
	}
	return
}

func (c *cOrder) GetOwnOrder(ctx context.Context, req *v1.GetOwnOrderReq) (res *v1.GetOwnOrderRes, err error) {
	res = &v1.GetOwnOrderRes{}
	userId := service.Session().GetUser(ctx).Id
	orders, err := service.Order().GetOrdersByUserId(ctx, userId)
	if err != nil {
		return
	}
	if orders == nil {
		res.Order = make([]*model.AdminResponseOrderForm, 0)
		return
	}
	for _, order := range orders {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Order = append(res.Order, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
	}

	return
}

func (c *cOrderManager) StartOrder(ctx context.Context, req *v1.StartOrderReq) (res *v1.StartOrderRes, err error) {
	res = &v1.StartOrderRes{}
	err = service.Order().StartOrder(ctx, req.OrderCode)
	if err != nil {
		return
	}
	return
}

func (c *cOrderManager) PayByCash(ctx context.Context, req *v1.PayOrderByCashReq) (res *v1.PayOrderByCashRes, err error) {
	res = &v1.PayOrderByCashRes{}
	err = service.Order().PayByCash(ctx, req.OrderCode)
	if err != nil {
		return
	}
	return
}

func (c *cOrder) GetOrderStatus(ctx context.Context, req *v1.GetOrdersByStatusReq) (res *v1.GetOrdersByStatusRes, err error) {
	res = &v1.GetOrdersByStatusRes{}
	orders, err := service.Order().GetOrderByStatus(ctx, req.Status)
	if err != nil {
		return
	}
	if orders == nil {
		res.Order = make([]*model.AdminResponseOrderForm, 0)
	}
	for _, order := range orders {
		place, err1 := service.Place().GetPlaceById(ctx, order.PlaceId)
		if err1 != nil {
			err = err1
		}
		res.Order = append(res.Order, &model.AdminResponseOrderForm{
			Order: order,
			Place: place,
		})
	}
	return
}
