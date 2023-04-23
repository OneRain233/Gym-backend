package v1

import (
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type FacilityReq struct {
	g.Meta `path:"/facility/facility" tags:"Facility" method:"post" summary:"Get All Facilities"`
	Limit  int `json:"limit"`
	Page   int `json:"page"`
}

type FacilityRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*model.FacilityEntity `json:"facility"`
}

type FacilitySearchReq struct {
	g.Meta `path:"/facility/search" tags:"Facility" method:"post" summary:"Get Facilities By searching tags"`
	Name   string `json:"name"`
	ID     int    `json:"id" `
}

type FacilitySearchRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*model.FacilityEntity `json:"facility"`
}

type AddFacilityReq struct {
	g.Meta      `path:"/facility/add" tags:"Facility" method:"post" summary:"Add Facility"`
	Name        string `json:"name" v:"required#Please input name"`
	Description string `json:"description" v:"required#Please input description"`
	Location    string `json:"location" v:"required#Please input location"`
	// TODO: images
	Image []string `json:"image" v:"required#Please input image"`
	Tags  []string `json:"tags" v:"required#Please input tag_ids"`
	Lat   float64  `json:"lat" v:"required#Please input lat"`
	Long  float64  `json:"long" v:"required#Please input long"`
}

type AddFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type AddFacilityPlaceReq struct {
	g.Meta      `path:"/facility/place/add" tags:"Facility" method:"post" summary:"Add Facility Place"`
	Name        string  `json:"name" v:"required#Please input name"`
	FacilityID  int     `json:"facility_id" v:"required#Please input facility_id"`
	Cost        float64 `json:"cost" v:"required#Please input cost"`
	Description string  `json:"description" v:"required#Please input description"`
}

type AddFacilityPlaceRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type ModifyFacilityPlaceReq struct {
	g.Meta      `path:"/facility/place/update" tags:"Facility" method:"post" summary:"Modify Facility Place"`
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
	g.Meta      `path:"/facility/update" tags:"Facility" method:"post" summary:"Modify Facility"`
	ID          int      `json:"id" v:"required#Please input id"`
	Name        string   `json:"name" v:"required#Please input name"`
	Description string   `json:"description" v:"required#Please input description"`
	Location    string   `json:"location" v:"required#Please input location"`
	Image       []string `json:"image" v:"required#Please input image"`
	Tags        []string `json:"tags" v:"required#Please input tag_ids"`
	Lat         float64  `json:"lat" v:"required#Please input lat"`
	Long        float64  `json:"long" v:"required#Please input long"`
}

type ModifyFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type FacilityDetailReq struct {
	g.Meta `path:"/facility/detail" tags:"Facility" method:"post" summary:"Get Facility Detail"`
	ID     int `json:"id" v:"required#Please input id"`
}

type FacilityDetailRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility *model.FacilityEntity `json:"facility"`
}

type OccupiedFacilityPlaceReq struct {
	g.Meta  `path:"/facility/place/occupied" tags:"Facility" method:"post" summary:"Get Occupied Facility Place"`
	PlaceId int `json:"place_id" v:"required#Please input place_id"`
}

type OccupiedFacilityPlaceRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Occupied []*model.OccupiedFacilityPlace `json:"occupied"`
}

type DeleteFacilityReq struct {
	g.Meta `path:"/facility/delete" tags:"Facility" method:"post" summary:"Delete Facility"`
	ID     int `json:"id" v:"required#Please input id"`
}

type DeleteFacilityRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type DeleteFacilityPlaceReq struct {
	g.Meta `path:"/facility/place/delete" tags:"Facility" method:"post" summary:"Delete Facility Place"`
	ID     int `json:"id" v:"required#Please input id"`
}

type DeleteFacilityPlaceRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type GetFacilityByTagReq struct {
	g.Meta `path:"/facility/tag" tags:"Facility" method:"post" summary:"Get Facility By Tag"`
	TagId  int `json:"tag_id" v:"required#Please input tag_id"`
}

type GetFacilityByTagRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*model.FacilityEntity `json:"facility"`
}

type GetFacilityByTagNameReq struct {
	g.Meta `path:"/facility/tag/name" tags:"Facility" method:"post" summary:"Get Facility By Tag Name"`
	Name   string `json:"name" v:"required#Please input name"`
}

type GetFacilityByTagNameRes struct {
	g.Meta   `mime:"application/json" example:"string"`
	Facility []*model.FacilityEntity `json:"facility"`
}
