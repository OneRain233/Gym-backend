package model

type CreateOrderForm struct {
	UserId    int
	PlaceId   int
	StartTime string
	EndTime   string
	Amount    float64
}
