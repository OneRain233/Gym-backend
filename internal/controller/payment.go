package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Payment = cPayment{}
var PaymentAdmin = cPaymentAdmin{}

type cPayment struct{}
type cPaymentAdmin struct{}

func (c *cPayment) CreatePayment(ctx context.Context, req *v1.CreatePaymentReq) (res *v1.CreatePaymentRes, err error) {
	res = &v1.CreatePaymentRes{}
	form := model.CreatePaymentForm{
		OrderCode:   req.OrderCode,
		PaymentType: req.PaymentType,
		CardId:      req.CardId,
	}
	resp, err := service.Payment().CreatePayment(ctx, &form)
	if err != nil {
		return
	}
	res.PaymentCode = resp.PaymentCode
	res.OrderCode = resp.OrderCode
	res.PaymentType = resp.PaymentType
	res.Amount = resp.Amount
	res.Status = resp.Status
	return
}
