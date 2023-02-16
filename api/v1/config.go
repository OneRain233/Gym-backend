package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetConfigReq struct {
	g.Meta `path:"/config/get" tags:"Config" method:"get" summary:"Get Config"`
}

type GetConfigRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Config []*entity.Config `json:"config"`
}

type UpdateConfigReq struct {
	g.Meta `path:"/config/update" tags:"Config" method:"post" summary:"Update Config"`
	Key    string `json:"key" v:"required#Please input key"`
	Value  string `json:"value" v:"required#Please input value"`
}

type UpdateConfigRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
