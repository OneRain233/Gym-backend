package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetCourseListReq struct {
	g.Meta `path:"/course/list" tags:"Course" method:"get" summary:"Get Course List"`
}

type GetCourseListRes struct {
	g.Meta  `mime:"application/json" example:"string"`
	Courses []*entity.Course `json:"courses"`
}

type AddCourseReq struct {
	g.Meta      `path:"/course/add" tags:"Course" method:"post" mime:"application/json" summary:"Add Course"`
	Title       string `json:"title" v:"required#Please input course title"`
	Description string `json:"description" v:"required#Please input course description"`
	Video       string `json:"video" v:"required#Please input course video"`
	Image       string `json:"image" v:"required#Please input course image"`
}

type AddCourseRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type GetCourseByIdReq struct {
	g.Meta `path:"/course/get" tags:"Course" method:"get" summary:"Get Course By Id"`
	Id     int `json:"id" v:"required#Please input course id"`
}

type GetCourseByIdRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Course *entity.Course `json:"course"`
}
