package model

import "Gym-backend/internal/model/entity"

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

type RefundPaymentForm struct {
	Order           *entity.Order   `json:"order"`
	OriginalPayment *entity.Payment `json:"originalPayment"`
}

type SubscriptionPaymentForm struct {
	OrderCode   string `json:"orderCode"`
	PaymentType int    `json:"paymentType"`
	CardId      int    `json:"cardId"`
}
