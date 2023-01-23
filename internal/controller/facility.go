package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

var Facility = cFacility{}

type cFacility struct{}

func (c *cFacility) GetAllFacility(ctx context.Context, req *v1.FacilityReq) (res *v1.FacilityRes, err error) {
	res = &v1.FacilityRes{}
	res.Facility, err = service.Facility().GetFacilityList()
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
		aFacility, err = service.Facility().GetFacilityById(id)
		if err != nil {
			return
		}
		facilities = append(facilities, aFacility)
	} else if name != "" {
		facilities, err = service.Facility().GetFacilityBySearch(name)
	} else {
		facilities, err = service.Facility().GetFacilityList()
	}

	res.Facility = facilities
	return
}
