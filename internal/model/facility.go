package model

import "Gym-backend/internal/model/entity"

type FacilityEntity struct {
	Facility *entity.Facility        `json:"facility"`
	Places   []*entity.FacilityPlace `json:"places"`
}

type AddFacilityForm struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Images      []string `json:"images"`
	Tags        []string `json:"tags"`
}

type ModifyFacilityForm struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Images      []string `json:"images"`
	Tags        []string `json:"tags"`
}

type AddFacilityPlaceForm struct {
	FacilityId  int     `json:"facilityId"`
	Name        string  `json:"name"`
	Cost        float64 `json:"cost"`
	Description string  `json:"description"`
}

type ModifyFacilityPlaceForm struct {
	Id          int     `json:"id"`
	FacilityId  int     `json:"facilityId"`
	Name        string  `json:"name"`
	Cost        float64 `json:"cost"`
	Description string  `json:"description"`
}

type OccupiedFacilityPlace struct {
	PlaceId   int    `json:"placeId"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
