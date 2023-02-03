package v1

import (
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type FacilityReq struct {
	g.Meta `path:"/facility" tags:"Facility" method:"get" summary:"Get All Facilities"`
}

type FacilityRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*model.FacilityEntity
}

type FacilitySearchReq struct {
	g.Meta `path:"/facility_search" tags:"Facility" method:"post" summary:"Get Facilities By searching tags"`
	Name   string `json:"name"`
	ID     int    `json:"id" `
}

type FacilitySearchRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*model.FacilityEntity
}

type AddFacilityReq struct {
	g.Meta      `path:"/add_facility" tags:"Facility" method:"post" summary:"Add Facility"`
	Name        string `json:"name" v:"required#Please input name"`
	Description string `json:"description" v:"required#Please input description"`
	Location    string `json:"location" v:"required#Please input location"`
	// TODO: images
	Image []string `json:"image" v:"required#Please input image"`
}

type AddFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type AddFacilityPlaceReq struct {
	g.Meta      `path:"/add_facility_place" tags:"Facility" method:"post" summary:"Add Facility Place"`
	Name        string  `json:"name" v:"required#Please input name"`
	FacilityID  int     `json:"facility_id" v:"required#Please input facility_id"`
	Cost        float64 `json:"cost" v:"required#Please input cost"`
	Description string  `json:"description" v:"required#Please input description"`
}

type AddFacilityPlaceRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type ModifyFacilityPlaceReq struct {
	g.Meta      `path:"/modify_facility_place" tags:"Facility" method:"post" summary:"Modify Facility Place"`
	ID          int     `json:"id" v:"required#Please input id"`
	Name        string  `json:"name" v:"required#Please input name"`
	FacilityID  int     `json:"facility_id" v:"required#Please input facility_id"`
	Cost        float64 `json:"cost" v:"required#Please input cost"`
	Description string  `json:"description" v:"required#Please input description"`
}

type ModifyFacilityPlaceRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type ModifyFacilityReq struct {
	g.Meta      `path:"/modify_facility" tags:"Facility" method:"post" summary:"Modify Facility"`
	ID          int      `json:"id" v:"required#Please input id"`
	Name        string   `json:"name" v:"required#Please input name"`
	Description string   `json:"description" v:"required#Please input description"`
	Location    string   `json:"location" v:"required#Please input location"`
	Image       []string `json:"image" v:"required#Please input image"`
	//Cost        float64 `json:"cost" v:"required#Please input cost"`
}

type ModifyFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type FacilityDetailReq struct {
	g.Meta `path:"/facility_detail" tags:"Facility" method:"post" summary:"Get Facility Detail"`
	ID     int `json:"id" v:"required#Please input id"`
}

type FacilityDetailRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility *model.FacilityEntity
}
