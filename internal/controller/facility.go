package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Facility = cFacility{}
var FacilityAdmin = cFacilityAdmin{}
var FacilityManager = cFacilityManager{}

type cFacility struct{}
type cFacilityAdmin struct{}
type cFacilityManager struct{}

func (c *cFacility) GetAllFacility(ctx context.Context, req *v1.FacilityReq) (res *v1.FacilityRes, err error) {
	res = &v1.FacilityRes{}
	pagination := &model.Pagination{
		Limit: req.Limit,
		Page:  req.Page,
	}
	res.Facility, err = service.Facility().GetFacilityList(ctx, pagination)
	if err != nil {
		return
	}
	return
}
func (c *cFacility) GetFacilityBySearching(ctx context.Context, req *v1.FacilitySearchReq) (res *v1.FacilitySearchRes, err error) {
	res = &v1.FacilitySearchRes{}
	name := req.Name
	id := req.ID
	pagination := &model.Pagination{
		Limit: 10,
		Page:  1,
	}
	var facilities []*model.FacilityEntity
	var aFacility *model.FacilityEntity
	if id != 0 {
		aFacility, err = service.Facility().GetFacilityById(ctx, id)
		if err != nil {
			return
		}
		facilities = append(facilities, aFacility)
	} else if name != "" {
		facilities, err = service.Facility().GetFacilityBySearch(ctx, name)
	} else {
		facilities, err = service.Facility().GetFacilityList(ctx, pagination)
	}

	res.Facility = facilities
	return
}

func (c *cFacility) GetFacilityByTag(ctx context.Context, req *v1.GetFacilityByTagReq) (res *v1.GetFacilityByTagRes, err error) {
	res = &v1.GetFacilityByTagRes{}
	res.Facility, err = service.Facility().GetFacilityByTagId(ctx, req.TagId)
	if err != nil {
		return
	}
	return
}

func (c *cFacility) GetFacilityByTagName(ctx context.Context, req *v1.GetFacilityByTagNameReq) (res *v1.GetFacilityByTagNameRes, err error) {
	res = &v1.GetFacilityByTagNameRes{}
	res.Facility, err = service.Facility().GetFacilityByTagName(ctx, req.Name)
	if err != nil {
		return
	}
	if res.Facility == nil {
		res.Facility = make([]*model.FacilityEntity, 0)
	}
	return
}

func (c *cFacilityAdmin) AddFacility(ctx context.Context, req *v1.AddFacilityReq) (res *v1.AddFacilityRes, err error) {
	res = &v1.AddFacilityRes{}
	form := model.AddFacilityForm{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Images:      req.Image,
		Tags:        req.Tags,
		Lat:         req.Lat,
		Long:        req.Long,
	}

	err = service.Facility().AddFacility(ctx, &form)
	if err != nil {
		return
	}
	return
}

func (c *cFacilityAdmin) ModifyFacility(ctx context.Context, req *v1.ModifyFacilityReq) (res *v1.ModifyFacilityRes, err error) {
	res = &v1.ModifyFacilityRes{}
	form := model.ModifyFacilityForm{
		Id:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Images:      req.Image,
		Tags:        req.Tags,
		Lat:         req.Lat,
		Long:        req.Long,
	}
	err = service.Facility().ModifyFacility(ctx, &form)
	if err != nil {
		return
	}
	return
}

func (c *cFacility) GetFacilityDetail(ctx context.Context, req *v1.FacilityDetailReq) (res *v1.FacilityDetailRes, err error) {
	res = &v1.FacilityDetailRes{}
	facility, err := service.Facility().GetFacilityById(ctx, req.ID)
	res.Facility = facility
	if err != nil {
		return
	}
	return
}

func (c *cFacilityAdmin) AddFacilityPlace(ctx context.Context, req *v1.AddFacilityPlaceReq) (res *v1.AddFacilityPlaceRes, err error) {
	res = &v1.AddFacilityPlaceRes{}
	form := model.AddFacilityPlaceForm{
		FacilityId:  req.FacilityID,
		Name:        req.Name,
		Cost:        req.Cost,
		Description: req.Description,
	}
	err = service.Facility().AddFacilityPlace(ctx, &form)
	if err != nil {
		return
	}
	return
}

func (c *cFacilityAdmin) ModifyFacilityPlace(ctx context.Context, req *v1.ModifyFacilityPlaceReq) (res *v1.ModifyFacilityPlaceRes, err error) {
	res = &v1.ModifyFacilityPlaceRes{}
	form := model.ModifyFacilityPlaceForm{
		Id:          req.ID,
		FacilityId:  req.FacilityID,
		Name:        req.Name,
		Cost:        req.Cost,
		Description: req.Description,
	}
	err = service.Facility().ModifyFacilityPlace(ctx, &form)
	if err != nil {
		return
	}
	return
}

func (c *cFacility) GetOccupiedFacilityPlaces(ctx context.Context, req *v1.OccupiedFacilityPlaceReq) (res *v1.OccupiedFacilityPlaceRes, err error) {
	res = &v1.OccupiedFacilityPlaceRes{}
	places, err := service.Facility().GetOccupiedFacilityPlaces(ctx, req.PlaceId)
	if places == nil {
		res.Occupied = make([]*model.OccupiedFacilityPlace, 0)
	} else {
		res.Occupied = places
	}
	if err != nil {
		return
	}
	return
}

func (c *cFacilityAdmin) DeleteFacility(ctx context.Context, req *v1.DeleteFacilityReq) (res *v1.DeleteFacilityRes, err error) {
	res = &v1.DeleteFacilityRes{}
	err = service.Facility().DeleteFacility(ctx, req.ID)
	if err != nil {
		return
	}
	return
}

func (c *cFacilityAdmin) DeleteFacilityPlace(ctx context.Context, req *v1.DeleteFacilityPlaceReq) (res *v1.DeleteFacilityPlaceRes, err error) {
	res = &v1.DeleteFacilityPlaceRes{}
	err = service.Place().DeletePlaceById(ctx, req.ID)
	if err != nil {
		return
	}
	return
}
