package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var File = cFile{}

type cFile struct{}

func (c *cFile) UploadFile(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	res = &v1.FileUploadRes{}
	output, err := service.File().UploadFile(ctx, model.FileUploadForm{
		File: req.File,
		Type: "common", // common file
	})
	if err != nil {
		return
	}
	res.Url = output.URL
	res.Name = output.Name
	return
}

func (c *cFile) UploadAvatar(ctx context.Context, req *v1.AvatarUploadReq) (res *v1.AvatarUploadRes, err error) {
	res = &v1.AvatarUploadRes{}
	output, err := service.File().UploadFile(ctx, model.FileUploadForm{
		File: req.File,
		Type: "avatar", // avatar file
	})
	if err != nil {
		return
	}
	res.Url = output.URL
	res.Name = output.Name
	return
}

func (c *cFile) UploadFacilityImage(ctx context.Context, req *v1.FacilityImageUploadReq) (res *v1.FacilityImageUploadRes, err error) {
	res = &v1.FacilityImageUploadRes{}
	output, err := service.File().UploadFile(ctx, model.FileUploadForm{
		File: req.File,
		Type: "facility", // facility image file
	})
	if err != nil {
		return
	}
	res.Url = output.URL
	res.Name = output.Name
	return
}
