// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model"
	"context"
)

type (
	IVisualize interface {
		GetFacilityWeeklyUsage(ctx context.Context) (res []*model.FacilityWeeklyUsage, err error)
		GetFacilityMonthlyUsage(ctx context.Context) (res []*model.FacilityMonthlyUsage, err error)
		GetWeeklyIncomeForAYear(ctx context.Context) (res []*model.WeeklyIncome, err error)
		GetDailyIncome(ctx context.Context, timeRange *model.TimeRange) (res []*model.DailyIncome, err error)
	}
)

var (
	localVisualize IVisualize
)

func Visualize() IVisualize {
	if localVisualize == nil {
		panic("implement not found for interface IVisualize, forgot register?")
	}
	return localVisualize
}

func RegisterVisualize(i IVisualize) {
	localVisualize = i
}
