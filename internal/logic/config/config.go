package config

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sConfig struct {
}

func init() {
	service.RegisterConfig(New())
}

func New() *sConfig {
	return &sConfig{}
}

func (c *sConfig) InitConfigToCache(ctx context.Context) error {
	err := service.Cache().Delete(ctx, "config")
	if err != nil {
		return gerror.New("Config make cache failed")
	}
	var configs []*entity.Config
	err = dao.Config.Ctx(ctx).Scan(&configs)
	if err != nil {
		return gerror.New("Config make cache failed")
	}
	for _, config := range configs {
		err = service.Cache().Push(ctx, "config", config)
		if err != nil {
			continue
			//return gerror.New("Config make cache failed")
		}
		err = service.Cache().Set(ctx, config.Key, config, 0)
		if err != nil {
			continue
			//return gerror.New("Config make cache failed")
		}
	}
	return nil
}

func (c *sConfig) GetConfig(ctx context.Context) (res []*entity.Config, err error) {
	// check cache
	// if cache exists, return cache
	// else, get from db and set cache
	// if cache not exists, get from db and set cache
	var cacheRes interface{}
	cacheRes, err = service.Cache().GetList(ctx, "config")
	if cacheRes != nil {
		//g.Log().Info(ctx, "cache exists", cacheRes)
		// convert interface to []*entity.Config

		for _, v := range cacheRes.([]interface{}) {
			//g.Log().Info(ctx, "cache exists", v)
			// parse json to entity
			var config *entity.Config
			err = gconv.Struct(v, &config)
			if err != nil {
				return
			}
			res = append(res, config)
		}
		return
	}
	//g.Log().Info(ctx, "cache not exists")
	// no cache
	err = dao.Config.Ctx(ctx).Scan(&res)
	if err != nil {
		return
	}
	// set cache
	_ = c.InitConfigToCache(ctx)
	return
}

func (c *sConfig) GetConfigByKey(ctx context.Context, key string) (res *entity.Config, err error) {
	// check cache
	// if cache exists, return cache
	// else, get from db and set cache
	var cacheRes interface{}
	cacheRes, err = service.Cache().Get(ctx, key)
	if cacheRes != nil {
		//g.Log().Info(ctx, "cache exists", cacheRes)
		// convert interface to *entity.Config
		err = gconv.Struct(cacheRes, &res)
		if err != nil {
			return
		}
		return
	}

	err = dao.Config.Ctx(ctx).Where("key", key).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (c *sConfig) UpdateConfig(ctx context.Context, config *model.Config) (err error) {
	// check if the config exists
	// if not, create it
	var configEntity *entity.Config
	configEntity, err = c.GetConfigByKey(ctx, config.Key)
	if err != nil {
		return
	}
	if configEntity == nil {
		err = c.CreateConfig(ctx, &model.Config{
			Key:   config.Key,
			Value: config.Value,
		})
		if err != nil {
			return
		}
		return
	} else {
		_, err = dao.Config.Ctx(ctx).Where("id", configEntity.Id).Data(config).Update()
		if err != nil {
			return
		}
	}
	// update cache
	_ = c.InitConfigToCache(ctx)
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
	// update cache
	_ = c.InitConfigToCache(ctx)
	return
}
