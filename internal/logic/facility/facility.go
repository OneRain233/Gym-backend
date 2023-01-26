package facility

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
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
func (s *sFacility) GetFacilityList(ctx context.Context) (res []*model.FacilityEntity, err error) {
	// get all facility first
	var facilities []*entity.Facility
	err = dao.Facility.Ctx(ctx).Scan(&facilities)
	if err != nil {
		return
	}
	// get all facility image
	res, err = s.FetchFacilityImages(ctx, facilities)
	return
}

// GetFacilityById gets the facility by id.
func (s *sFacility) GetFacilityById(ctx context.Context, id int) (res *model.FacilityEntity, err error) {
	res = &model.FacilityEntity{}
	err = dao.Facility.Ctx(ctx).Where("id", id).Scan(&res.Facility)
	if err != nil {
		return
	}
	err = dao.FacilityImage.Ctx(ctx).Where("facilityID", id).Scan(&res.Images)
	if err != nil {
		return
	}
	return
}

// GetFacilityByName gets the facility by name.
func (s *sFacility) GetFacilityByName(ctx context.Context, name string) (res *model.FacilityEntity, err error) {
	res = &model.FacilityEntity{}
	err = dao.Facility.Ctx(ctx).Where("name", name).Scan(&res)
	return
}

// GetFacilityBySearch gets the facility by search in name.
func (s *sFacility) GetFacilityBySearch(ctx context.Context, search string) (res []*model.FacilityEntity, err error) {
	//err = dao.Facility.Ctx(ctx).Where("name like ?", "%"+search+"%").Scan(&res)
	res = []*model.FacilityEntity{}
	// get all facility first
	var facilities []*entity.Facility
	err = dao.Facility.Ctx(ctx).Where("name like ?", "%"+search+"%").Scan(&facilities)
	if err != nil {
		return
	}
	// get all facility image
	res, err = s.FetchFacilityImages(ctx, facilities)
	if err != nil {
		return
	}

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

func (s *sFacility) GetFacilityImages(ctx context.Context, id int) (res []*entity.FacilityImage, err error) {
	err = dao.FacilityImage.Ctx(ctx).Where("facilityID", id).Scan(&res)
	return
}

func (s *sFacility) FetchFacilityImages(ctx context.Context, facilities []*entity.Facility) (res []*model.FacilityEntity, err error) {
	res = []*model.FacilityEntity{}

	var facilityIds []int
	for _, facility := range facilities {
		facilityIds = append(facilityIds, facility.Id)
	}
	var facilityImages []*entity.FacilityImage
	err = dao.FacilityImage.Ctx(ctx).Where("facilityID in (?)", facilityIds).Scan(&facilityImages)
	if err != nil {
		return
	}
	// map facility id to facility image
	facilityImageMap := make(map[int][]*entity.FacilityImage)
	for _, facilityImage := range facilityImages {
		facilityImageMap[facilityImage.FacilityID] = append(facilityImageMap[facilityImage.FacilityID], facilityImage)
	}

	// map facility to facility entity
	for _, facility := range facilities {
		res = append(res, &model.FacilityEntity{
			Facility: facility,
			Images:   facilityImageMap[facility.Id],
		})
	}

	return
}
