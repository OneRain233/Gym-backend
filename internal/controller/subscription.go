package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/consts"
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

func (c *cSubscription) CancelSubscription(ctx context.Context, req *v1.CancelSubscriptionReq) (res *v1.CancelSubscriptionRes, err error) {
	res = &v1.CancelSubscriptionRes{}
	err = service.Subscription().CancelSubscription(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cSubscriptionAdmin) AdminCancelSubscription(ctx context.Context, req *v1.AdminCancelSubscriptionReq) (res *v1.AdminCancelSubscriptionRes, err error) {
	res = &v1.AdminCancelSubscriptionRes{}
	err = service.Subscription().CancelSubscriptionByUserId(ctx, req.UserId)
	if err != nil {
		return
	}
	return
}

func (c *cSubscriptionAdmin) AdminGetSubscription(ctx context.Context, req *v1.AdminGetSubscriptionReq) (res *v1.AdminGetSubscriptionRes, err error) {
	res = &v1.AdminGetSubscriptionRes{}
	res.Subscription, err = service.Subscription().GetAllSubscriptions(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cSubscription) AdminCancelSubscriptionById(ctx context.Context, req *v1.AdminCancelSubscriptionByIdReq) (res *v1.AdminCancelSubscriptionByIdRes, err error) {
	res = &v1.AdminCancelSubscriptionByIdRes{}
	err = service.Subscription().UpdateSubscriptionStatus(ctx, req.Id, consts.SubscriptionStatusCancel)
	if err != nil {
		return
	}
	return
}

func (c *cSubscriptionAdmin) AdminUpdateSubscriptionType(ctx context.Context, req *v1.AdminUpdateSubscriptionTypeReq) (res *v1.AdminUpdateSubscriptionTypeRes, err error) {
	res = &v1.AdminUpdateSubscriptionTypeRes{}
	form := &model.UpdateSubscriptionTypeForm{
		Id:     req.Id,
		Amount: req.Amount,
		Days:   req.Days,
	}
	err = service.Subscription().UpdateSubscriptionType(ctx, form)
	if err != nil {
		return
	}
	return
}
