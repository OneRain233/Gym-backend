package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Config = cConfig{}

type cConfig struct {
}

func (c *cConfig) GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error) {
	res = &v1.GetConfigRes{}
	res.Config, err = service.Config().GetConfig(ctx)
	if err != nil {
		return
	}
	return
}

func (c *cConfig) UpdateConfig(ctx context.Context, req *v1.UpdateConfigReq) (res *v1.UpdateConfigRes, err error) {
	res = &v1.UpdateConfigRes{}
	configs := req.Configs
	for _, config := range configs {
		err = service.Config().UpdateConfig(ctx, &model.Config{
			Key:   config.Key,
			Value: config.Value,
			Type:  config.Type,
		})
		if err != nil {
			return
		}
	}
	return
}
