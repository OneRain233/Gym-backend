package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type FacilityReq struct {
	g.Meta `path:"/facility" tags:"Facility" method:"get" summary:"Get All Facilities"`
}

type FacilityRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*entity.Facility
}

type FacilitySearchReq struct {
	g.Meta `path:"/facility_search" tags:"Facility" method:"post" summary:"Get Facilities By searching tags"`
	Name   string
	ID     int
}

type FacilitySearchRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*entity.Facility
}
