package v1

import "github.com/gogf/gf/v2/frame/g"

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
