package facility

import (
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"

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
func (s *sFacility) GetFacilityList() (res []*entity.Facility, err error) {
	err = g.DB().Model("facility").Scan(&res)
	return
}

// GetFacilityById gets the facility by id.
func (s *sFacility) GetFacilityById(id int) (res *entity.Facility, err error) {
	err = g.DB().Model("facility").Where("id", id).Scan(&res)
	return
}

// GetFacilityByName gets the facility by name.
func (s *sFacility) GetFacilityByName(name string) (res *entity.Facility, err error) {
	err = g.DB().Model("facility").Where("name", name).Scan(&res)
	return
}

// GetFacilityBySearch gets the facility by search in name.
func (s *sFacility) GetFacilityBySearch(search string) (res []*entity.Facility, err error) {
	err = g.DB().Model("facility").WhereLike("name", search).Scan(&res)
	return
}
