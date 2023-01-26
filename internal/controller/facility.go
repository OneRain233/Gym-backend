package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

var Facility = cFacility{}
var FacilityAdmin = cFacilityAdmin{}

type cFacility struct{}
type cFacilityAdmin struct{}

func (c *cFacility) GetAllFacility(ctx context.Context, req *v1.FacilityReq) (res *v1.FacilityRes, err error) {
	res = &v1.FacilityRes{}
	res.Facility, err = service.Facility().GetFacilityList(ctx)
	if err != nil {
		return
	}
	return
}
func (c *cFacility) GetFacilityBySearching(ctx context.Context, req *v1.FacilitySearchReq) (res *v1.FacilitySearchRes, err error) {
	res = &v1.FacilitySearchRes{}
	name := req.Name
	id := req.ID
	var facilities []*entity.Facility
	var aFacility *entity.Facility
	if id != 0 {
		aFacility, err = service.Facility().GetFacilityById(ctx, id)
		if err != nil {
			return
		}
		facilities = append(facilities, aFacility)
	} else if name != "" {
		facilities, err = service.Facility().GetFacilityBySearch(ctx, name)
	} else {
		facilities, err = service.Facility().GetFacilityList(ctx)
	}

	res.Facility = facilities
	return
}

func (c *cFacilityAdmin) AddFacility(ctx context.Context, req *v1.AddFacilityReq) (res *v1.AddFacilityRes, err error) {
	res = &v1.AddFacilityRes{}
	facility := &entity.Facility{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Cost:        req.Cost,
	}

	err = service.Facility().AddFacility(ctx, facility)
	if err != nil {
		return
	}
	return
}

func (c *cFacilityAdmin) ModifyFacility(ctx context.Context, req *v1.ModifyFacilityReq) (res *v1.ModifyFacilityRes, err error) {
	res = &v1.ModifyFacilityRes{}
	facility := &entity.Facility{
		Id:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Cost:        req.Cost,
	}
	err = service.Facility().ModifyFacility(ctx, facility)
	if err != nil {
		return
	}
	return
}
