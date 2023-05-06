package model

import "Gym-backend/internal/model/entity"

type CreateOrderForm struct {
	UserId    int
	PlaceId   int
	StartTime string
	EndTime   string
}

type ResponseOrderForm struct {
	Amount    float64
	PlaceId   int
	StartTime string
	EndTime   string
	OrderCode string
}

type RegularOrderResponseForm struct {
	Amount    float64
	OrderCode string
	Orders    []*ResponseOrderForm
}

type RegularOrderInfoResponseForm struct {
	Amount    float64
	OrderCode string
	Orders    []*entity.Order
}

type CreateReceiptForm struct {
	OrderCode   string
	OrderId     int
	ReceiptPath string
}

type AdminResponseOrderForm struct {
	Order *entity.Order         `json:"order"`
	Place *entity.FacilityPlace `json:"place"`
}

type CreateRegularOrderFormWeekly struct {
	UserId           int
	PlaceId          int
	SessionStartTime string
	SessionEndTime   string
	StartDay         string
	WeekCount        int
}

type CreateRegularOrderFormDaily struct {
	UserId           int
	PlaceId          int
	SessionStartTime string
	SessionEndTime   string
	Weekday          int
}
