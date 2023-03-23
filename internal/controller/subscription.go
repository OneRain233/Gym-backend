package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Subscription = cSubscription{}
var SubscriptionAdmin = cSubscriptionAdmin{}

type cSubscription struct{}
type cSubscriptionAdmin struct{}

func (c *cSubscription) GetOwnSubscription(ctx context.Context, req *v1.GetOwnSubscriptionReq) (res *v1.GetOwnSubscriptionRes, err error) {
	res = &v1.GetOwnSubscriptionRes{}
	user := service.Session().GetUser(ctx)
	res.Subscription, err = service.Subscription().GetSubscriptionListByUserId(ctx, user.Id)
	if err != nil {
		return
	}
	return
}

func (c *cSubscription) CreateSubscription(ctx context.Context, req *v1.CreateSubscriptionReq) (res *v1.CreateSubscriptionRes, err error) {
	subscriptionType := req.Type
	paymentType := req.PaymentType
	days := 0
	switch subscriptionType {
	case 1:
		days = 30
	case 2:
		days = 365
	}
	user := service.Session().GetUser(ctx)
	form := &model.CreateSubscriptionForm{
		UserId:      user.Id,
		Days:        days,
		PaymentType: paymentType,
		CardId:      req.CardId,
	}
	err = service.Subscription().CreateSubscription(ctx, form)

	if err != nil {
		return
	}
	return
}
