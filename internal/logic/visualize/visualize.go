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

func (v *sVisualize) GetWeeklyIncomeForAYear(ctx context.Context) (res []*model.WeeklyIncome, err error) {
	// used to draw bar chart
	return
}

func (v *sVisualize) GetDailyIncome(ctx context.Context, timeRange *model.TimeRange) (res []*model.DailyIncome, err error) {
	// used to draw line chart
	startDate := timeRange.StartDate
	endDate := timeRange.EndDate
	for startDate.Before(endDate) {
		income, err := service.Order().GetDailyOrderIncome(ctx, startDate)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.DailyIncome{
			Date:   startDate.FormatNew("Y-m-d"),
			Income: income,
		})
		startDate = startDate.AddDate(0, 0, 1)
	}
	return
}

func (v *sVisualize) GetWeeklyIncome(ctx context.Context, timeRange *model.TimeRange) (res []*model.WeeklyIncome, err error) {
	// used to draw line chart
	startDate := timeRange.StartDate.StartOfWeek()
	endDate := timeRange.EndDate.EndOfWeek()
	for startDate.Before(endDate) {
		income, err := service.Order().GetWeeklyOrderIncome(ctx, startDate)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.WeeklyIncome{
			Week:   startDate.WeeksOfYear(),
			Income: income,
		})
		startDate = startDate.AddDate(0, 0, 7)
	}
	return
}
