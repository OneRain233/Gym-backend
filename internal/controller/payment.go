package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

var Payment = cPayment{}
var PaymentAdmin = cPaymentAdmin{}

type cPayment struct{}
type cPaymentAdmin struct{}

func (c *cPayment) CreatePayment(ctx context.Context, req *v1.CreatePaymentReq) (res *v1.CreatePaymentRes, err error) {
	res = &v1.CreatePaymentRes{}
	length := len(req.OrderCode)
	fmt.Println(length)
	if len(req.OrderCode) == 17 {
		g.Log().Info(ctx, "order code is 17")
		form := model.CreatePaymentForm{
			OrderCode:   req.OrderCode,
			PaymentType: req.PaymentType,
			CardId:      req.CardId,
		}
		resp, err1 := service.Payment().CreatePayment(ctx, &form)
		if err1 != nil {
			return
		}
		res.PaymentCode = resp.PaymentCode
		res.OrderCode = resp.OrderCode
		res.PaymentType = resp.PaymentType
		res.Amount = resp.Amount
		res.Status = resp.Status
		return
	} else if len(req.OrderCode) == 18 {
		g.Log().Info(ctx, "order code is 18")
		form := model.CreatePaymentForm{
			OrderCode:   req.OrderCode,
			PaymentType: req.PaymentType,
			CardId:      req.CardId,
		}
		resp, err1 := service.Payment().CreatePaymentForRegularOrder(ctx, &form)
		if err1 != nil {
			return
		}
		res.PaymentCode = resp.PaymentCode
		res.OrderCode = resp.OrderCode
		res.PaymentType = resp.PaymentType
		res.Amount = resp.Amount
		res.Status = resp.Status
		return
	} else {
		err = gerror.New("order code not valid")
		return
	}
}

func (c *cPaymentAdmin) GetPaymentByUserId(ctx context.Context, req *v1.GetPaymentByUserIdReq) (res *v1.GetPaymentByUserIdRes, err error) {
	res = &v1.GetPaymentByUserIdRes{}
	fmt.Println(req.UserId)
	if req.UserId == 0 {
		fmt.Println("get all payment")
		res.Payments, err = service.Payment().GetAllPayment(ctx)
		if err != nil {
			return
		}
		return
	} else {
		fmt.Println("get payment by user id")
		res.Payments, err = service.Payment().GetPaymentByUserId(ctx, req.UserId)
		if err != nil {
			return
		}
		return
	}
}

func (c *cPayment) GetOwnPayment(ctx context.Context, req *v1.GetOwnPaymentReq) (res *v1.GetOwnPaymentRes, err error) {
	res = &v1.GetOwnPaymentRes{}
	user := service.Session().GetUser(ctx)
	if user == nil {
		err = gerror.New("user not found")
		return
	}
	res.Payments, err = service.Payment().GetPaymentByUserId(ctx, user.Id)
	if err != nil {
		return
	}
	return
}
