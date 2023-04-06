package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var SubscriptionUnauthorized = cSubscriptionUnauthorized{}
var Subscription = cSubscription{}
var SubscriptionAdmin = cSubscriptionAdmin{}

type cSubscription struct{}
type cSubscriptionAdmin struct{}
type cSubscriptionUnauthorized struct{}

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
	user := service.Session().GetUser(ctx)
	form := &model.CreateSubscriptionForm{
		UserId:      user.Id,
		Type:        subscriptionType,
		PaymentType: paymentType,
		CardId:      req.CardId,
	}
	err = service.Subscription().CreateSubscription(ctx, form)

	if err != nil {
		return
	}
	return
}

func (c *cSubscriptionUnauthorized) GetSubscriptionType(ctx context.Context, req *v1.GetSubscriptionTypeReq) (res *v1.GetSubscriptionTypeRes, err error) {
	res = &v1.GetSubscriptionTypeRes{}
	res.SubscriptionType, err = service.Subscription().GetTypesOfSubscription(ctx)
	if err != nil {
		return
	}
	return
}