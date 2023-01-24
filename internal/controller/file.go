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
		Type: req.Type,
		Name: req.Name,
	})
	if err != nil {
		return
	}
	res.Data.Url = output.URL
	res.Data.Name = output.Name
	return
}
