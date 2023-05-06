package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

var Visualize = cVisualize{}

type cVisualize struct{}

func (c *cVisualize) GetDailyIncome(ctx context.Context, req *v1.GetDailyIncomeReq) (res *v1.GetDailyIncomeRes, err error) {
	res = &v1.GetDailyIncomeRes{}
	startDate := gtime.NewFromStr(req.StartDate)
	endDate := gtime.NewFromStr(req.EndDate)
	timeRange := &model.TimeRange{
		StartDate: startDate,
		EndDate:   endDate,
	}
	res.Incomes, err = service.Visualize().GetDailyIncome(ctx, timeRange)
	if err != nil {
		return
	}
	return
}

func (c *cVisualize) GetWeeklyIncome(ctx context.Context, req *v1.GetWeeklyIncomeReq) (res *v1.GetWeeklyIncomeRes, err error) {
	res = &v1.GetWeeklyIncomeRes{}
	startDate := gtime.NewFromStr(req.StartDate)
	endDate := gtime.NewFromStr(req.EndDate)
	timeRange := &model.TimeRange{
		StartDate: startDate,
		EndDate:   endDate,
	}
	res.Incomes, err = service.Visualize().GetWeeklyIncome(ctx, timeRange)
	if err != nil {
		return
	}
	return
}
