package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateSubscriptionReq struct {
	g.Meta      `path:"/subscription/add" tags:"Subscription" method:"post" summary:"Add Subscription"`
	Type        int `json:"type" v:"required#Please input type"`
	PaymentType int `json:"payment_type" v:"required#Please input payment_type"`
	CardId      int `json:"card_id"`
}

type CreateSubscriptionRes struct{}

type GetOwnSubscriptionReq struct {
	g.Meta `path:"/subscription/get" tags:"Subscription" method:"post" summary:"Get Own Subscription"`
}

type GetOwnSubscriptionRes struct {
	Subscription []*entity.Subscription `json:"subscription"`
}

type GetSubscriptionTypeReq struct {
	g.Meta `path:"/subscription/type" tags:"Subscription" method:"post" summary:"Get Subscription Type"`
}

type GetSubscriptionTypeRes struct {
	SubscriptionType []*entity.SubscriptionType `json:"subscription_type"`
}

type CancelSubscriptionReq struct {
	g.Meta `path:"/subscription/cancel" tags:"Subscription" method:"post" summary:"Cancel Subscription"`
}

type CancelSubscriptionRes struct{}

type AdminCancelSubscriptionReq struct {
	g.Meta `path:"/subscription/admin/cancel" tags:"Subscription" method:"post" summary:"Admin Cancel Subscription"`
	UserId int `json:"user_id" v:"required#Please input user_id"`
}

type AdminCancelSubscriptionRes struct{}

type AdminGetSubscriptionReq struct {
	g.Meta `path:"/subscription/admin/get" tags:"Subscription" method:"post" summary:"Admin Get Subscription"`
}

type AdminGetSubscriptionRes struct {
	Subscription []*entity.Subscription `json:"subscription"`
}

type AdminCancelSubscriptionByIdReq struct {
	g.Meta `path:"/subscription/admin/cancel-by-id" tags:"Subscription" method:"post" summary:"Admin Cancel Subscription By Id"`
	Id     int `json:"id" v:"required#Please input id"`
}

type AdminCancelSubscriptionByIdRes struct{}

type AdminUpdateSubscriptionTypeReq struct {
	g.Meta `path:"/subscription/admin/update-type" tags:"Subscription" method:"post" summary:"Admin Update Subscription Type"`
	Id     int     `json:"id" v:"required#Please input id"`
	Name   string  `json:"name" v:"required#Please input name"`
	Days   int     `json:"days" v:"required#Please input days"`
	Amount float64 `json:"amount" v:"required#Please input amount"`
}

type AdminUpdateSubscriptionTypeRes struct{}
