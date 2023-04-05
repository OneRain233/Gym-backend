package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/service"
	"context"
)

var Tag = cTag{}
var TagAdmin = cTagAdmin{}

type cTag struct{}
type cTagAdmin struct{}

func (c *cTag) GetAllTags(ctx context.Context, req *v1.GetAllTagsReq) (res *v1.GetAllTagsRes, err error) {
	res = &v1.GetAllTagsRes{}
	res.Tags, err = service.Tag().GetAllTags(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cTagAdmin) AddTag(ctx context.Context, req *v1.AddTagReq) (res *v1.AddTagRes, err error) {
	res = &v1.AddTagRes{}
	err = service.Tag().CreateNewTag(ctx, req.Name)
	if err != nil {
		return
	}
	return
}
