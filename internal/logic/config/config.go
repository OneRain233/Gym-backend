package config

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sConfig struct {
}

func init() {
	service.RegisterConfig(New())
}

func New() *sConfig {
	return &sConfig{}
}

func (c *sConfig) GetConfig(ctx context.Context) (res []*entity.Config, err error) {
	err = dao.Config.Ctx(ctx).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (c *sConfig) GetConfigByKey(ctx context.Context, key string) (res *entity.Config, err error) {
	err = dao.Config.Ctx(ctx).Where("key", key).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (c *sConfig) UpdateConfig(ctx context.Context, config *entity.Config) (err error) {
	_, err = dao.Config.Ctx(ctx).Where("id", config.Id).Update(config)
	if err != nil {
		return
	}
	return
}

func (c *sConfig) DeleteConfig(ctx context.Context, id int) (err error) {
	_, err = dao.Config.Ctx(ctx).Where("id", id).Delete()
	return
}

func (c *sConfig) CreateConfig(ctx context.Context, config *model.Config) (err error) {
	var configEntity entity.Config
	configEntity.Key = config.Key
	configEntity.Value = config.Value

	_, err = dao.Config.Ctx(ctx).Insert(&configEntity)
	if err != nil {
		return
	}

	return
}
