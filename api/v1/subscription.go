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
