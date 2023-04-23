package course

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sCourse struct {
}

func init() {
	service.RegisterCourse(New())
}

func New() *sCourse {
	return &sCourse{}
}

func (s *sCourse) GetCourseList(ctx context.Context) (res []*entity.Course, err error) {
	err = dao.Course.Ctx(ctx).Scan(&res)
	if err != nil {
		return
	}
	if len(res) == 0 {
		res = []*entity.Course{}
	}
	return
}

func (s *sCourse) AddCourse(ctx context.Context, form *model.AddCourseForm) (err error) {
	course := &entity.Course{
		Title:       form.Title,
		Description: form.Description,
		Video:       form.Video,
		Image:       form.Image,
	}
	_, err = dao.Course.Ctx(ctx).Insert(course)
	if err != nil {
		return
	}
	return
}

func (s *sCourse) GetCourseById(ctx context.Context, id int) (res *entity.Course, err error) {
	res = &entity.Course{}
	err = dao.Course.Ctx(ctx).Where("id", id).Scan(&res)
	if err != nil {
		return
	}
	return
}
