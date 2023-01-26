package model

import "Gym-backend/internal/model/entity"

type FacilityEntity struct {
	Facility *entity.Facility        `json:"facility"`
	Images   []*entity.FacilityImage `json:"images"`
}
