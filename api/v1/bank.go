package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddBankReq struct {
	g.Meta `path:"/bank/add" method:"post" mime:"application/json" tags:"Bank" summary:"Add bank"`
	Name   string `json:"name" v:"required#Please input bank name"`
}

type AddBankRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type UpdateBankReq struct {
	g.Meta `path:"/bank/update" method:"post" mime:"application/json" tags:"Bank" summary:"Update bank"`
	Id     int    `json:"id" v:"required#Please input bank id"`
	Name   string `json:"name" v:"required#Please input bank name"`
}

type UpdateBankRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type GetBanksReq struct {
	g.Meta `path:"/bank/banks" method:"get" tags:"Bank" summary:"Get banks"`
}

type GetBanksRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   []*entity.Bank `json:"data"`
}

type DeleteBankReq struct {
	g.Meta `path:"/bank/delete" method:"post" mime:"application/json" tags:"Bank" summary:"Delete bank"`
	Id     int `json:"id" v:"required#Please input bank id"`
}

type DeleteBankRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
