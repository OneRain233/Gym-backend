package v1

import (
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllTagsReq struct {
	g.Meta `path:"/tag/all" method:"get" tags:"Tag" summary:"Get all tags"`
}

type GetAllTagsRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Tags   []*entity.Tag `json:"tags"`
}

type AddTagReq struct {
	g.Meta `path:"/tag/add" method:"post" tags:"Tag" summary:"Add tag"`
	Name   string `json:"name" v:"required#Please input name"`
}

type AddTagRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type DeleteTagReq struct {
	g.Meta `path:"/tag/delete" method:"post" tags:"Tag" summary:"Delete tag"`
	ID     int `json:"id" v:"required#Please input id"`
}

type DeleteTagRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
