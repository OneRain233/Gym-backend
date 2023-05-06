package v1

import (
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDailyIncomeReq struct {
	g.Meta    `path:"/visualize/income/daily" method:"post" tags:"Visualize" summary:"Get daily income"`
	StartDate string `json:"start_date" v:"required#Please input start date"`
	EndDate   string `json:"end_date" v:"required#Please input end date"`
}

type GetDailyIncomeRes struct {
	g.Meta  `mime:"application/json" example:"string"`
	Incomes []*model.DailyIncome `json:"incomes"`
}

type GetWeeklyIncomeReq struct {
	g.Meta    `path:"/visualize/income/weekly" method:"post" tags:"Visualize" summary:"Get weekly income"`
	StartDate string `json:"start_date" v:"required#Please input start date"`
	EndDate   string `json:"end_date" v:"required#Please input end date"`
}

type GetWeeklyIncomeRes struct {
	g.Meta  `mime:"application/json" example:"string"`
	Incomes []*model.WeeklyIncome `json:"incomes"`
}
