package model

type CreatePaymentForm struct {
	OrderCode   string `json:"orderCode"`
	PaymentType int    `json:"paymentType"`
	CardId      int    `json:"cardId"`
}

type ResponsePaymentForm struct {
	OrderCode   string  `json:"orderCode"`
	PaymentType int     `json:"paymentType"`
	Amount      float64 `json:"amount"`
	PaymentCode string  `json:"paymentCode"`
	Status      int     `json:"status"`
}
