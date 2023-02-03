package model

type CreatePaymentForm struct {
	WalletId    int     `json:"walletId"`
	OrderId     int     `json:"orderId"`
	Amount      float64 `json:"amount"`
	PaymentType int     `json:"paymentType"`
}
