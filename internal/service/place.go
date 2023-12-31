// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model/entity"
	"context"
)

type (
	IPlace interface {
		GetPlaceById(ctx context.Context, id int) (res *entity.FacilityPlace, err error)
		DeletePlaceById(ctx context.Context, id int) error
	}
)

var (
	localPlace IPlace
)

func Place() IPlace {
	if localPlace == nil {
		panic("implement not found for interface IPlace, forgot register?")
	}
	return localPlace
}

func RegisterPlace(i IPlace) {
	localPlace = i
}
