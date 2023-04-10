package tag

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sTag struct{}

func init() {
	service.RegisterTag(New())
}

func New() *sTag {
	return &sTag{}
}

func (s *sTag) CreateNewTag(ctx context.Context, name string) (err error) {
	tag := &entity.Tag{
		Name: name,
	}
	_, err = dao.Tag.Ctx(ctx).Insert(tag)
	return
}

func (s *sTag) DeleteTag(ctx context.Context, tagId int) (err error) {
	_, err = dao.Tag.Ctx(ctx).Where("id", tagId).Delete()
	return
}

func (s *sTag) GetAllTags(ctx context.Context) (res []*entity.Tag, err error) {
	res = make([]*entity.Tag, 0)
	err = dao.Tag.Ctx(ctx).Scan(&res)
	return
}

func (s *sTag) GetTagById(ctx context.Context, id int) (res *entity.Tag, err error) {
	res = &entity.Tag{}
	err = dao.Tag.Ctx(ctx).Where("id", id).Scan(res)
	return
}

func (s *sTag) GetTagByName(ctx context.Context, name string) (res *entity.Tag, err error) {
	res = &entity.Tag{}
	err = dao.Tag.Ctx(ctx).Where("name", name).Scan(res)
	return
}
