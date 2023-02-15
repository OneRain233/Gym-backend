package facility

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"strings"

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

func (s *sFacility) ProcessImage(images string) []string {
	return strings.Split(images, ",")
}

// GetFacilityList gets the facility list.
func (s *sFacility) GetFacilityList(ctx context.Context) (res []*model.FacilityEntity, err error) {
	// get all facility first
	var facilities []*entity.Facility
	err = dao.Facility.Ctx(ctx).Scan(&facilities)
	if err != nil {
		return
	}

	// map the places to facility
	for _, facility := range facilities {
		var place []*entity.FacilityPlace
		err = dao.FacilityPlace.Ctx(ctx).Where("facility_id", facility.Id).Scan(&place)
		if err != nil {
			return
		}
		if place == nil {
			place = []*entity.FacilityPlace{}
		}
		res = append(res, &model.FacilityEntity{
			Facility: facility,
			Places:   place,
		})
	}
	return
}

// GetFacilityById gets the facility by id.
func (s *sFacility) GetFacilityById(ctx context.Context, id int) (res *model.FacilityEntity, err error) {
	res = &model.FacilityEntity{}
	err = dao.Facility.Ctx(ctx).Where("id", id).Scan(&res.Facility)
	if err != nil {
		return
	}
	var place []*entity.FacilityPlace
	err = dao.FacilityPlace.Ctx(ctx).Where("id", id).Scan(&place)
	if err != nil {
		return
	}
	if place == nil {
		place = []*entity.FacilityPlace{}
	}
	res.Places = place
	return
}

// GetFacilityByName gets the facility by name.
func (s *sFacility) GetFacilityByName(ctx context.Context, name string) (res *model.FacilityEntity, err error) {
	res = &model.FacilityEntity{}
	err = dao.Facility.Ctx(ctx).Where("name", name).Scan(&res.Facility)
	if err != nil {
		return
	}
	var place []*entity.FacilityPlace
	err = dao.FacilityPlace.Ctx(ctx).Where("id", res.Facility.Id).Scan(&place)
	if err != nil {
		return
	}
	if place == nil {
		place = []*entity.FacilityPlace{}
	}
	res.Places = place
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
	// map the places to facility
	for _, facility := range facilities {
		var place []*entity.FacilityPlace
		err = dao.FacilityPlace.Ctx(ctx).Where("id", facility.Id).Scan(&place)
		if err != nil {
			return
		}
		if place == nil {
			place = []*entity.FacilityPlace{}
		}
		res = append(res, &model.FacilityEntity{
			Facility: facility,
			Places:   place,
		})
	}
	return
}

// AddFacility adds the facility.
func (s *sFacility) AddFacility(ctx context.Context, input *model.AddFacilityForm) (err error) {
	// check if all the fields are filled
	err = s.ValidateAddFacility(ctx, input)
	if err != nil {
		return
	}
	// convert images to string like "image1,image2,image3"
	var images string
	for _, image := range input.Images {
		images += image + ","
	}

	var facility = &entity.Facility{
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
		Images:      images,
	}

	_, err = g.DB().Model("facility").Data(facility).Insert()
	return
}

// ModifyFacility modifies the facility.
func (s *sFacility) ModifyFacility(ctx context.Context, input *model.ModifyFacilityForm) (err error) {
	// check if all the fields are filled
	if input.Id == 0 {
		err = gerror.New("Id is empty")
		return
	}
	// merge previous data
	var facility *entity.Facility
	err = dao.Facility.Ctx(ctx).Where("id", input.Id).Scan(&facility)
	if err != nil {
		return
	}
	if facility == nil {
		err = gerror.New("facility not found")
		return
	}
	if input.Name != "" {
		facility.Name = input.Name
	}
	if input.Description != "" {
		facility.Description = input.Description
	}
	if input.Location != "" {
		facility.Location = input.Location
	}
	if input.Images != nil {
		var images string
		for _, image := range input.Images {
			images += image + ","
		}
		facility.Images = images
	}
	_, err = g.DB().Model("facility").Data(facility).Where("id", input.Id).Update()
	if err != nil {
		return
	}
	return
}

// ValidateAddFacility validates the facility.
func (s *sFacility) ValidateAddFacility(ctx context.Context, facility *model.AddFacilityForm) (err error) {
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

	if facility.Images == nil {
		err = gerror.New("Images is empty")
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

func (s *sFacility) AddFacilityPlace(ctx context.Context, input *model.AddFacilityPlaceForm) (err error) {
	if input.FacilityId == 0 {
		err = gerror.New("FacilityId is empty")
		return
	}
	if input.Name == "" {
		err = gerror.New("Name is empty")
		return
	}
	if input.Cost == 0 {
		err = gerror.New("Cost is empty")
		return
	}

	// check if the facility is exist
	var facility *entity.Facility
	err = dao.Facility.Ctx(ctx).Where("id", input.FacilityId).Scan(&facility)
	if err != nil {
		return
	}
	if facility == nil {
		err = gerror.New("The facility is not exist")
		return
	}

	var place = &entity.FacilityPlace{
		FacilityId:  input.FacilityId,
		Name:        input.Name,
		Cost:        input.Cost,
		Description: input.Description,
	}

	_, err = g.DB().Model("facility_place").Data(place).Insert()
	return
}

func (s *sFacility) ModifyFacilityPlace(ctx context.Context, input *model.ModifyFacilityPlaceForm) (err error) {
	if input.Id == 0 {
		err = gerror.New("Id is empty")
		return
	}
	if input.Name == "" {
		err = gerror.New("Name is empty")
		return
	}
	if input.Cost == 0 {
		err = gerror.New("Cost is empty")
		return
	}

	// check if the facility is exist
	var place *entity.FacilityPlace
	err = dao.FacilityPlace.Ctx(ctx).Where("id", input.Id).Scan(&place)
	if err != nil {
		return
	}
	if place == nil {
		err = gerror.New("The place is not exist")
		return
	}

	if input.Name != "" {
		place.Name = input.Name
	}
	if input.Cost != 0 {
		place.Cost = input.Cost
	}
	if input.Description != "" {
		place.Description = input.Description
	}
	if input.FacilityId != 0 {
		// check if the facility is exist
		var facility *entity.Facility
		err = dao.Facility.Ctx(ctx).Where("id", input.FacilityId).Scan(&facility)
		if err != nil {
			return
		}
		if facility == nil {
			err = gerror.New("The facility is not exist")
			return
		}

		place.FacilityId = input.FacilityId
	}

	_, err = g.DB().Model("facility_place").Data(place).Where("id", input.Id).Update()
	return
}

func (s *sFacility) GetOccupiedFacilityPlaces(ctx context.Context, placeId int) (res []*model.OccupiedFacilityPlace, err error) {
	// check if place exist
	var place *entity.FacilityPlace
	err = dao.FacilityPlace.Ctx(ctx).Where("id", placeId).Scan(&place)
	if err != nil {
		return
	}
	if place == nil {
		err = gerror.New("place not found")
		return
	}
	// get all the orders
	var orders []*entity.Order
	err = dao.Order.Ctx(ctx).Where("place_id", placeId).Scan(&orders)
	if err != nil {
		return
	}
	if orders == nil {
		return
	}
	for _, order := range orders {
		occupied := model.OccupiedFacilityPlace{
			PlaceId:   placeId,
			StartTime: order.StartTime.String(),
			EndTime:   order.EndTime.String(),
		}
		res = append(res, &occupied)
	}
	return
}
