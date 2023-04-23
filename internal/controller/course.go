package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Course = cCourse{}
var CourseAdmin = cCourseAdmin{}

type cCourse struct{}
type cCourseAdmin struct{}

func (c *cCourse) GetCourseList(ctx context.Context, req *v1.GetCourseListReq) (res *v1.GetCourseListRes, err error) {
	res = &v1.GetCourseListRes{}
	res.Courses, err = service.Course().GetCourseList(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cCourse) GetCourseById(ctx context.Context, req *v1.GetCourseByIdReq) (res *v1.GetCourseByIdRes, err error) {
	res = &v1.GetCourseByIdRes{}
	res.Course, err = service.Course().GetCourseById(ctx, req.Id)
	if err != nil {
		return
	}
	return
}

func (c *cCourse) AddCourse(ctx context.Context, req *v1.AddCourseReq) (res *v1.AddCourseRes, err error) {
	res = &v1.AddCourseRes{}
	form := &model.AddCourseForm{
		Title:       req.Title,
		Description: req.Description,
		Video:       req.Video,
		Image:       req.Image,
	}
	err = service.Course().AddCourse(ctx, form)
	if err != nil {
		return
	}
	return
}
