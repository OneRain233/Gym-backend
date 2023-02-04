package session

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

// SetUser set user content to session
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return service.BizCtx().Get(ctx).Session.Set(consts.UserKey, user)
}

// GetUser return user info
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	c := service.BizCtx().Get(ctx)
	if c != nil {
		v, _ := c.Session.Get(consts.UserKey)
		if !v.IsNil() {
			var user *entity.User
			_ = v.Structs(&user)
			return user
		}
	}
	return &entity.User{}
}

func (s *sSession) RemoveUser(ctx context.Context) error {
	c := service.BizCtx().Get(ctx)
	if c != nil {
		return c.Session.Remove(consts.UserKey)
	}
	return nil
}
