package facility

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
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

func (s *sFacility) FetchTags(ctx context.Context, facility *entity.Facility) (tags []string, err error) {
	if facility == nil {
		return nil, gerror.New("facility is nil")
	}
	if facility.Tags == "" || facility == nil {
		return make([]string, 0), nil
	}
	tagIdsStr := strings.Split(facility.Tags, ",")
	for _, tagIdStr := range tagIdsStr {
		tagId, err := strconv.Atoi(tagIdStr)
		if err != nil {
			return nil, gerror.New("tag id is not valid")
		}
		var tag *entity.Tag
		tag, err = service.Tag().GetTagById(ctx, tagId)
		if err != nil {
			return nil, gerror.New("tag id is not valid")
		}
		tags = append(tags, tag.Name)
	}
	return
}

// GetFacilityList gets the facility list.
func (s *sFacility) GetFacilityList(ctx context.Context, pagination *model.Pagination) (res []*model.FacilityEntity, err error) {
	// get all facility first
	var facilities []*entity.Facility
	if pagination.Limit == 0 || pagination.Page == 0 {
		err = dao.Facility.Ctx(ctx).Scan(&facilities)
	} else {
		offset := (pagination.Page - 1) * pagination.Limit
		err = dao.Facility.Ctx(ctx).Limit(pagination.Limit).Offset(offset).Scan(&facilities)
	}
	if err != nil {
		return
	}
	// get tags
	for _, facility := range facilities {
		tags, err := s.FetchTags(ctx, facility)
		if err != nil {
			return nil, err
		}
		facility.Tags = strings.Join(tags, ",")
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
		rate, _ := service.Evaluation().GetFacilityScore(ctx, facility.Id)
		res = append(res, &model.FacilityEntity{
			Facility: facility,
			Places:   place,
			Rate:     float64(rate),
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
	tags, err := s.FetchTags(ctx, res.Facility)
	if err != nil {
		return nil, err
	}
	res.Facility.Tags = strings.Join(tags, ",")
	var place []*entity.FacilityPlace
	err = dao.FacilityPlace.Ctx(ctx).Where("facility_id", id).Scan(&place)
	if err != nil {
		return
	}
	if place == nil {
		place = []*entity.FacilityPlace{}
	}
	res.Places = place
	rate, err := service.Evaluation().GetFacilityScore(ctx, id)
	if err != nil {
		return
	}
	res.Rate = float64(rate)
	return
}

// GetFacilityByName gets the facility by name.
func (s *sFacility) GetFacilityByName(ctx context.Context, name string) (res *model.FacilityEntity, err error) {
	res = &model.FacilityEntity{}
	err = dao.Facility.Ctx(ctx).Where("name", name).Scan(&res.Facility)
	if err != nil {
		return
	}
	tags, err := s.FetchTags(ctx, res.Facility)
	if err != nil {
		return nil, err
	}
	res.Facility.Tags = strings.Join(tags, ",")
	var place []*entity.FacilityPlace
	err = dao.FacilityPlace.Ctx(ctx).Where("id", res.Facility.Id).Scan(&place)
	if err != nil {
		return
	}
	if place == nil {
		place = []*entity.FacilityPlace{}
	}
	res.Places = place
	rate, err := service.Evaluation().GetFacilityScore(ctx, res.Facility.Id)
	if err != nil {
		return
	}
	res.Rate = float64(rate)
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
		if tags, err := s.FetchTags(ctx, facility); err != nil {
			return nil, err
		} else {
			facility.Tags = strings.Join(tags, ",")
		}

		var place []*entity.FacilityPlace
		err = dao.FacilityPlace.Ctx(ctx).Where("id", facility.Id).Scan(&place)
		if err != nil {
			return
		}
		if place == nil {
			place = []*entity.FacilityPlace{}
		}
		rate, _ := service.Evaluation().GetFacilityScore(ctx, facility.Id)
		res = append(res, &model.FacilityEntity{
			Facility: facility,
			Places:   place,
			Rate:     float64(rate),
		})
	}
	return
}

// GetFacilityByTagId gets the facility by tag.
func (s *sFacility) GetFacilityByTagId(ctx context.Context, tagId int) (res []*model.FacilityEntity, err error) {
	res = []*model.FacilityEntity{}
	// get all facility first
	var facilities []*entity.Facility
	err = dao.Facility.Ctx(ctx).Scan(&facilities)
	if err != nil {
		return
	}
	// map the places to facility
	for _, facility := range facilities {
		tags := strings.Split(facility.Tags, ",")
		for _, t := range tags {
			if t == strconv.Itoa(tagId) {
				if tags, err := s.FetchTags(ctx, facility); err != nil {
					return nil, err
				} else {
					facility.Tags = strings.Join(tags, ",")
				}

				var place []*entity.FacilityPlace
				err = dao.FacilityPlace.Ctx(ctx).Where("id", facility.Id).Scan(&place)
				if err != nil {
					return
				}
				if place == nil {
					place = []*entity.FacilityPlace{}
				}
				rate, _ := service.Evaluation().GetFacilityScore(ctx, facility.Id)
				res = append(res, &model.FacilityEntity{
					Facility: facility,
					Places:   place,
					Rate:     float64(rate),
				})
			}
		}
	}
	return
}

func (s *sFacility) GetFacilityByTagName(ctx context.Context, tagName string) (res []*model.FacilityEntity, err error) {
	// check if the tag is exist
	var tag *entity.Tag
	err = dao.Tag.Ctx(ctx).Where("name", tagName).Scan(&tag)
	if err != nil {
		return
	}
	if tag == nil {
		err = gerror.New("The tag is not exist")
		return
	}
	// get all the facilities
	var facilities []*entity.Facility
	err = dao.Facility.Ctx(ctx).Scan(&facilities)
	if err != nil {
		return
	}
	if facilities == nil {
		return
	}
	for _, facility := range facilities {
		tagIds := strings.Split(facility.Tags, ",")
		// find the tag
		for _, tagId := range tagIds {
			if tagId == strconv.Itoa(tag.Id) {
				if tags, err := s.FetchTags(ctx, facility); err != nil {
					return nil, err
				} else {
					facility.Tags = strings.Join(tags, ",")
				}
				var place []*entity.FacilityPlace
				err = dao.FacilityPlace.Ctx(ctx).Where("id", facility.Id).Scan(&place)
				if err != nil {
					return
				}
				if place == nil {
					place = []*entity.FacilityPlace{}
				}
				rate, _ := service.Evaluation().GetFacilityScore(ctx, facility.Id)
				res = append(res, &model.FacilityEntity{
					Facility: facility,
					Places:   place,
					Rate:     float64(rate),
				})
			}
		}
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
	//for _, image := range input.Images {
	//	images += image + ","
	//}
	images = strings.Join(input.Images, ",")

	var tags string
	var tagIds []string
	for _, tag := range input.Tags {
		// to string
		// check if the tag is exist
		var tagEntity *entity.Tag
		if tagEntity, err = service.Tag().GetTagByName(ctx, tag); err != nil {
			// create it
			if tagEntity == nil || tagEntity.Id == 0 {
				err = service.Tag().CreateNewTag(ctx, tag)
				if err != nil {
					return
				}
			}
		}
		tagEntity, err = service.Tag().GetTagByName(ctx, tag)
		if err != nil {
			return
		}
		// get the tag id
		tagIds = append(tagIds, strconv.Itoa(tagEntity.Id))

		//tagIdStr := strconv.Itoa(tag)
		//tagIds = append(tagIds, tagIdStr)
	}
	tags = strings.Join(tagIds, ",")

	var facility = &entity.Facility{
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
		Images:      images,
		Tags:        tags,
		Lat:         input.Lat,
		Long:        input.Long,
	}

	_, err = dao.Facility.Ctx(ctx).Insert(facility)
	return
}

// ValidateAddFacility validates the facility.
func (s *sFacility) ValidateAddFacility(ctx context.Context, facility *model.AddFacilityForm) (err error) {
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
	err = s.ValidateModifyFacility(ctx, input)
	if err != nil {
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
		images = strings.Join(input.Images, ",")
		facility.Images = images
	}
	if input.Lat != 0 {
		facility.Lat = input.Lat
	}
	if input.Long != 0 {
		facility.Long = input.Long
	}
	if input.Tags != nil {
		var tags string
		var tagIds []string
		for _, tag := range input.Tags {
			// to string
			// check if the tag is exist
			var tagEntity *entity.Tag
			if tagEntity, err = service.Tag().GetTagByName(ctx, tag); err != nil {
				// create it
				if tagEntity == nil || tagEntity.Id == 0 {
					err = service.Tag().CreateNewTag(ctx, tag)
					if err != nil {
						return
					}
				}
			}
			tagEntity, err = service.Tag().GetTagByName(ctx, tag)
			if err != nil {
				return
			}
			// get the tag id
			tagIds = append(tagIds, strconv.Itoa(tagEntity.Id))

		}
		tags = strings.Join(tagIds, ",")
		facility.Tags = tags
	}
	_, err = dao.Facility.Ctx(ctx).Where("id", input.Id).Update(facility)
	if err != nil {
		return
	}
	return
}

func (s *sFacility) ValidateModifyFacility(ctx context.Context, facility *model.ModifyFacilityForm) (err error) {
	cnt, err := dao.Facility.Ctx(ctx).Where("name", facility.Name).Count()
	if err != nil {
		return
	}
	if cnt > 1 {
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
	// check if the place is exist
	err = s.ValidateAddFacilityPlace(ctx, input.Name)
	if err != nil {
		return
	}
	var place = &entity.FacilityPlace{
		FacilityId:  input.FacilityId,
		Name:        input.Name,
		Cost:        input.Cost,
		Description: input.Description,
	}
	_, err = dao.FacilityPlace.Ctx(ctx).Insert(place)
	return
}

func (s *sFacility) ValidateAddFacilityPlace(ctx context.Context, facilityName string) (err error) {
	cnt, err := dao.FacilityPlace.Ctx(ctx).Where(dao.FacilityPlace.Columns().Name, facilityName).Count()
	if err != nil {
		return
	}
	if cnt > 0 {
		err = gerror.New("The facility already exists")
		return
	}
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

	_, err = dao.FacilityPlace.Ctx(ctx).Where("id", input.Id).Update(place)
	return
}

func (s *sFacility) GetOccupiedFacilityPlaces(ctx context.Context, placeId int) (res []*model.OccupiedFacilityPlace, err error) {
	// TODO: in a period of time

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
	// where place_id = placeId and (status = 1 or status = 2)
	err = dao.Order.Ctx(ctx).Where("place_id = ? and (status = ? or status = ?)", placeId, 1, 2).Scan(&orders)
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

func (s *sFacility) DeleteFacility(ctx context.Context, id int) (err error) {
	if id == 0 {
		err = gerror.New("Id is empty")
		return
	}
	_, err = dao.Facility.Ctx(ctx).Where("id", id).Delete()
	return
}
