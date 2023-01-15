package session

import (
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
func (s *sSession) SetUser(ctx context.Context) error {
	// TODO
	return nil
}

func (s *sSession) GetUser(ctx context.Context) error {
	// TODO
	return nil
}

func (s *sSession) RemoveUser(ctx context.Context) error {
	// TODO
	return nil
}
