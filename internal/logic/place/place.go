package place

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sPlace struct{}

func init() {
	service.RegisterPlace(New())
}

func New() *sPlace {
	return &sPlace{}
}

func (p *sPlace) GetPlaceById(ctx context.Context, id int) (res *entity.FacilityPlace, err error) {
	res = &entity.FacilityPlace{}
	err = dao.FacilityPlace.Ctx(ctx).Where("id", id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *sPlace) DeletePlaceById(ctx context.Context, id int) error {
	_, err := dao.FacilityPlace.Ctx(ctx).Where("id", id).Delete()
	return err
}
