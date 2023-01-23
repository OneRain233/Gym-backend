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

type AddFacilityReq struct {
	g.Meta      `path:"/add_facility" tags:"Facility" method:"post" summary:"Add Facility"`
	Name        string
	Description string
	Location    string
	Cost        float64
}

type AddFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type ModifyFacilityReq struct {
	g.Meta      `path:"/modify_facility" tags:"Facility" method:"post" summary:"Modify Facility"`
	ID          int
	Name        string
	Description string
	Location    string
	Cost        float64
}

type ModifyFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
