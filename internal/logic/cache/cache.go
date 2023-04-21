package cache

import (
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gredis"

	"github.com/gogf/gf/v2/frame/g"
)

type sCache struct{}

func New() *sCache {
	return &sCache{}
}

func init() {
	service.RegisterCache(New())
}

func (s *sCache) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := g.Redis().Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *sCache) GetList(ctx context.Context, key string) (interface{}, error) {
	value, err := g.Redis().LRange(ctx, key, 0, -1)
	if err != nil {
		return nil, err
	}
	return value.Interfaces(), nil
}

func (s *sCache) Set(ctx context.Context, key string, value interface{}, expire int) error {
	// expire == 0 means no expire time
	if expire == 0 {
		_, err := g.Redis().Set(ctx, key, value)
		if err != nil {
			return err
		}
	}

	expireSeconds := int64(expire)
	options := gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: &expireSeconds,
		},
	}
	_, err := g.Redis().Set(ctx, key, value, options)
	if err != nil {
		return err
	}
	return nil
}

func (s *sCache) Push(ctx context.Context, key string, value interface{}) error {
	// check if key exists
	// if exists, push to the end of the list
	// if not exists, create a new list and push to the end of the list

	// check if key exists
	exists, err := g.Redis().Exists(ctx, key)
	if err != nil {
		return err
	}
	if exists == 0 {
		_, err := g.Redis().LPush(ctx, key, value)
		if err != nil {
			return err
		}
		return nil
	} else if exists == 1 {
		_, err := g.Redis().RPush(ctx, key, value)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *sCache) Delete(ctx context.Context, key string) error {
	_, err := g.Redis().Del(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
