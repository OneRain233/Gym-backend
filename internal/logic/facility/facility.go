package facility

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sFacility struct{}

func init() {
	service.RegisterFacility(New())
}

func New() *sFacility {
	return &sFacility{}
}

// GetFacilityList gets the facility list.
func (s *sFacility) GetFacilityList(ctx context.Context) (res []*entity.Facility, err error) {
	err = dao.Facility.Ctx(ctx).Scan(&res)
	return
}

// GetFacilityById gets the facility by id.
func (s *sFacility) GetFacilityById(ctx context.Context, id int) (res *entity.Facility, err error) {
	err = dao.Facility.Ctx(ctx).Where("id", id).Scan(&res)
	return
}

// GetFacilityByName gets the facility by name.
func (s *sFacility) GetFacilityByName(ctx context.Context, name string) (res *entity.Facility, err error) {
	err = dao.Facility.Ctx(ctx).Where("name", name).Scan(&res)
	return
}

// GetFacilityBySearch gets the facility by search in name.
func (s *sFacility) GetFacilityBySearch(ctx context.Context, search string) (res []*entity.Facility, err error) {
	err = dao.Facility.Ctx(ctx).Where("name like ?", "%"+search+"%").Scan(&res)
	return
}

// AddFacility adds the facility.
func (s *sFacility) AddFacility(ctx context.Context, facility *entity.Facility) (err error) {
	// check if all the fields are filled
	err = s.ValidateFacility(ctx, facility)
	if err != nil {
		return
	}
	_, err = g.DB().Model("facility").Data(facility).Insert()
	return
}

// ModifyFacility modifies the facility.
func (s *sFacility) ModifyFacility(ctx context.Context, facility *entity.Facility) (err error) {
	// check if all the fields are filled
	err = s.ValidateFacility(ctx, facility)
	if err != nil {
		return
	}
	_, err = g.DB().Model("facility").Data(facility).Where("id", facility.Id).Update()
	return
}

// ValidateFacility validates the facility.
func (s *sFacility) ValidateFacility(ctx context.Context, facility *entity.Facility) (err error) {
	if facility.Name == "" {
		err = gerror.New("Name is empty")
		return
	}
	if facility.Description == "" {
		err = gerror.New("Description is empty")
		return
	}
	if facility.Location == "" {
		err = gerror.New("Location is empty")
		return
	}
	if facility.Cost == 0 {
		err = gerror.New("Cost is empty")
		return
	}

	cnt, err := dao.Facility.Ctx(ctx).Where("name", facility.Name).Count()
	if err != nil {
		return
	}
	if cnt > 0 {
		err = gerror.New("The facility already exists")
		return
	}
	return
}
