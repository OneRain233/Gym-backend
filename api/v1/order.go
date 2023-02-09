package v1

import (
	"Gym-backend/internal/model/entity"

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

type GetAllOrderReq struct {
	g.Meta `path:"/order/all" method:"post" tags:"Order" summary:"Get all order"`
}

type GetAllOrderRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*entity.Order `json:"order"`
}

type GetOrderByUserIdReq struct {
	g.Meta `path:"/order/user" method:"post" tags:"Order" summary:"Get order by user id"`
	UserId int `json:"userId" v:"required#Please input user id"`
}

type GetOrderByUserIdRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*entity.Order `json:"order"`
}

type GetOrderByPlaceIdReq struct {
	g.Meta  `path:"/order/place" method:"post" tags:"Order" summary:"Get order by place id"`
	PlaceId int `json:"placeId" v:"required#Please input place id"`
}

type GetOrderByPlaceIdRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*entity.Order `json:"order"`
}

type GetOrderByTimeReq struct {
	g.Meta    `path:"/order/time" method:"post" tags:"Order" summary:"Get order by time"`
	StartTime string `json:"startTime" v:"required#Please input start time"`
	EndTime   string `json:"endTime" v:"required#Please input end time"`
}

type GetOrderByTimeRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  []*entity.Order `json:"order"`
}

type GetOrderByOrderCodeReq struct {
	g.Meta    `path:"/order/code" method:"post" tags:"Order" summary:"Get order by order code"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type GetOrderByOrderCodeRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Order  *entity.Order `json:"order"`
}

type GetReceiptReq struct {
	g.Meta    `path:"/order/receipt" method:"post" tags:"Order" summary:"Get receipt by order code"`
	OrderCode string `json:"orderCode" v:"required#Please input order code"`
}

type GetReceiptRes struct {
	// file response
	g.Meta `mime:"application/octet-stream" example:"string"`
}
