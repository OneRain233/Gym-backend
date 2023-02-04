package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddBankReq struct {
	g.Meta `path:"/add_bank" method:"post" mime:"application/json" tags:"Bank" summary:"Add bank"`
	Name   string `json:"name" v:"required#Please input bank name"`
}

type AddBankRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type UpdateBankReq struct {
	g.Meta `path:"/update_bank" method:"post" mime:"application/json" tags:"Bank" summary:"Update bank"`
	Id     int    `json:"id" v:"required#Please input bank id"`
	Name   string `json:"name" v:"required#Please input bank name"`
}

type UpdateBankRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type GetBanksReq struct {
	g.Meta `path:"/banks" method:"get" tags:"Bank" summary:"Get banks"`
}

type GetBanksRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   []*entity.Bank
}
