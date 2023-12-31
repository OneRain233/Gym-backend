// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	ISubscription interface {
		InitSubscriptionSTypesToCache(ctx context.Context) error
		GetTypesOfSubscription(ctx context.Context) (res []*entity.SubscriptionType, err error)
		GetSubscriptionPermissionById(ctx context.Context, permissionId int) (res *entity.SubscriptionPermission, err error)
		GetSubscriptionTypeById(ctx context.Context, id int) (res *entity.SubscriptionType, err error)
		GetSubscriptionEndDayByUserId(ctx context.Context, userId int) (sub *entity.Subscription, res *gtime.Time, err error)
		GetUnFinishedSubscription(ctx context.Context, userId int) (res *entity.Subscription, err error)
		GetSubscriptionListByUserId(ctx context.Context, userId int) (res []*entity.Subscription, err error)
		GetSubscriptionByUserId(ctx context.Context, userId int) (res *entity.Subscription, err error)
		CreateSubscription(ctx context.Context, form *model.CreateSubscriptionForm) error
		UpdateSubscriptionStatus(ctx context.Context, subscriptionId int, status int) error
		GenerateOrderCode() string
		CancelSubscription(ctx context.Context) error
		CheckExpiredSubscription(ctx context.Context) error
		CancelSubscriptionByUserId(ctx context.Context, userId int) error
		GetAllSubscriptions(ctx context.Context) ([]*entity.Subscription, error)
		UpdateSubscriptionType(ctx context.Context, form *model.UpdateSubscriptionTypeForm) error
	}
)

var (
	localSubscription ISubscription
)

func Subscription() ISubscription {
	if localSubscription == nil {
		panic("implement not found for interface ISubscription, forgot register?")
	}
	return localSubscription
}

func RegisterSubscription(i ISubscription) {
	localSubscription = i
}
