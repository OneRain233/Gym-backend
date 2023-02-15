package visualize

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

type sVisualize struct{}

func init() {
	service.RegisterVisualize(New())
}

func New() *sVisualize {
	return &sVisualize{}
}

func (v *sVisualize) GetFacilityWeeklyUsage(ctx context.Context) (res []*model.FacilityWeeklyUsage, err error) {
	return
}

func (v *sVisualize) GetFacilityMonthlyUsage(ctx context.Context) (res []*model.FacilityMonthlyUsage, err error) {
	return
}
