package model

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

type CreateReceiptForm struct {
	OrderCode   string
	OrderId     int
	ReceiptPath string
}
