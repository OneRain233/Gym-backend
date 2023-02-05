package v1

import "github.com/gogf/gf/v2/frame/g"

type CreatePaymentReq struct {
	g.Meta      `path:"/order/payment" method:"post" mime:"application/json" tags:"Payment" summary:"Create payment"`
	OrderCode   string `json:"orderCode" v:"required#Please input order code"`
	PaymentType int    `json:"paymentType" v:"required#Please input payment type"`
	CardId      int    `json:"cardId"`
}

type CreatePaymentRes struct {
	g.Meta      `mime:"application/json" example:"string"`
	OrderCode   string  `json:"orderCode"`
	PaymentType int     `json:"paymentType"`
	Amount      float64 `json:"amount"`
	PaymentCode string  `json:"paymentCode"`
	Status      int     `json:"status"`
}
