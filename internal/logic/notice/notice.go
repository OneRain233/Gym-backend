package notice

import (
	"Gym-backend/internal/service"
)

type sNotice struct {
}

func New() *sNotice {
	return &sNotice{}
}

func init() {
	service.RegisterNotice(New())
}
