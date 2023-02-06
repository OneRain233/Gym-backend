package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePaymentReq struct {
	g.Meta      `path:"/order/payment/create" method:"post" mime:"application/json" tags:"Payment" summary:"Create payment"`
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

type GetPaymentByPaymentCodeReq struct {
	g.Meta      `path:"/order/payment/order" method:"post" mime:"application/json" tags:"Payment" summary:"Get payment by payment code"`
	PaymentCode string `json:"paymentCode" v:"required#Please input payment code"`
}

type GetPaymentByPaymentCodeRes struct {
	g.Meta  `mime:"application/json" example:"string"`
	Payment *entity.Payment `json:"payment"`
}

type GetPaymentByUserIdReq struct {
	g.Meta `path:"/order/payment/user" method:"post" mime:"application/json" tags:"Payment" summary:"Get payment by userId"`
	UserId int `json:"userId"`
}

type GetPaymentByUserIdRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Payments []*entity.Payment `json:"payments"`
}

type GetOwnPaymentReq struct {
	g.Meta `path:"/order/payment/own" method:"post" mime:"application/json" tags:"Payment" summary:"Get own payment"`
}

type GetOwnPaymentRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Payments []*entity.Payment `json:"payments"`
}
