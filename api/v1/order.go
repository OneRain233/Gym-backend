package v1

import (
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateOrderReq struct {
	g.Meta    `path:"/order/create" method:"post" tags:"Order" summary:"Create order"`
	PlaceId   int    `json:"placeId" v:"required#Please input place id"`
	StartTime string `json:"startTime" v:"required#Please input start time"`
	EndTime   string `json:"endTime" v:"required#Please input end time"`
}

type CreateOrderRes struct {
	g.Meta    `mime:"application/json" example:"string"`
	OrderCode string  `json:"orderCode"`
	Amount    float64 `json:"amount"`
	PlaceId   int     `json:"placeId"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
}

type CreateRegularOrderReq struct {
	g.Meta           `path:"/order/create-regular" method:"post" tags:"Order" summary:"Create regular order"`
	PlaceId          int    `json:"placeId" v:"required#Please input place id"`
	SessionStartTime string `json:"session_start_time" v:"required#Please input session start time"`
	SessionEndTime   string `json:"session_end_time" v:"required#Please input session end time"`
	StartDay         string `json:"start_day" v:"required#Please input start day"`
	WeekCount        int    `json:"week_count" v:"required#Please input week count"`
}

type CreateRegularOrderRes struct {
	g.Meta    `mime:"application/json" example:"string"`
	OrderCode string                          `json:"orderCode"`
	Orders    *model.RegularOrderResponseForm `json:"orders"`
}

type GetRegularOrderInfoReq struct {
	g.Meta    `path:"/order/regular-info" method:"post" tags:"Order" summary:"Get regular order info"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type GetRegularOrderInfoRes struct {
	g.Meta    `mime:"application/json" example:"string"`
	OrderCode string                          `json:"orderCode"`
	Amount    float64                         `json:"amount"`
	Orders    []*model.AdminResponseOrderForm `json:"orders"`
}

type GetAllOrderReq struct {
	g.Meta `path:"/order/all" method:"post" tags:"Order" summary:"Get all order"`
	Limit  int `json:"limit"`
	Page   int `json:"page"`
}

type GetAllOrderRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*model.AdminResponseOrderForm `json:"order"`
}

type GetOrderByUserIdReq struct {
	g.Meta `path:"/order/user" method:"post" tags:"Order" summary:"Get order by user id"`
	UserId int `json:"userId" v:"required#Please input user id"`
}

type GetOrderByUserIdRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*model.AdminResponseOrderForm `json:"order"`
}

type GetOrderByPlaceIdReq struct {
	g.Meta  `path:"/order/place" method:"post" tags:"Order" summary:"Get order by place id"`
	PlaceId int `json:"placeId" v:"required#Please input place id"`
}

type GetOrderByPlaceIdRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*model.AdminResponseOrderForm `json:"order"`
}

type GetOrderByTimeReq struct {
	g.Meta    `path:"/order/time" method:"post" tags:"Order" summary:"Get order by time"`
	StartTime string `json:"startTime" v:"required#Please input start time"`
	EndTime   string `json:"endTime" v:"required#Please input end time"`
}

type GetOrderByTimeRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*model.AdminResponseOrderForm `json:"order"`
}

type GetOrderByOrderCodeReq struct {
	g.Meta    `path:"/order/code" method:"post" tags:"Order" summary:"Get order by order code"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type GetOrderByOrderCodeRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  *model.AdminResponseOrderForm `json:"order"`
}

type GetReceiptReq struct {
	g.Meta    `path:"/order/receipt" method:"post" tags:"Order" summary:"Get receipt by order code"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type GetReceiptRes struct {
	// file response
	g.Meta `mime:"application/octet-stream" example:"string"`
}

type RefundOrderReq struct {
	g.Meta    `path:"/order/refund" method:"post" tags:"Order" summary:"Refund order by order code"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type RefundOrderRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type GetOwnOrderReq struct {
	g.Meta `path:"/order/own" method:"post" tags:"Order" summary:"Get own order"`
}

type GetOwnOrderRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*model.AdminResponseOrderForm `json:"order"`
}

type StartOrderReq struct {
	g.Meta    `path:"/order/start" method:"post" tags:"Order" summary:"Start order by order code"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type StartOrderRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type GetOrdersByStatusReq struct {
	g.Meta `path:"/order/status" method:"post" tags:"Order" summary:"Get orders by status"`
	Status int `json:"status" v:"required#Please input status"`
}

type GetOrdersByStatusRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*model.AdminResponseOrderForm `json:"order"`
}

type PayOrderByCashReq struct {
	g.Meta    `path:"/order/pay-cash" method:"post" tags:"Order" summary:"Pay order by cash"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type PayOrderByCashRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
