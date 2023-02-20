package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetCreditReq struct {
	g.Meta `path:"/credit/get" tags:"Credit" method:"get" summary:"Get Credit"`
}

type GetCreditRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Credit *entity.Credit `json:"credit"`
}

type AdminGetCreditReq struct {
	g.Meta `path:"/credit/all" tags:"Credit" method:"get" summary:"Admin Get Credit"`
}

type AdminGetCreditRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Credit []*entity.Credit `json:"credit"`
}
