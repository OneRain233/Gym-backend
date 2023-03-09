package model

import "github.com/gogf/gf/v2/os/gtime"

type FacilityWeeklyUsage struct {
	FacilityName string `json:"facilityName"`
	Usage        int    `json:"usage"`
}

type MultiFacilityWeeklyUsage struct {
	Facilities []FacilityWeeklyUsage `json:"facilities"`
}

type FacilityMonthlyUsage struct {
	FacilityName string `json:"facilityName"`
	Usage        int    `json:"usage"`
}

type MultiFacilityMonthlyUsage struct {
	Facilities []FacilityMonthlyUsage `json:"facilities"`
}

type WeeklyIncome struct {
	Week   int     `json:"week"`
	Income float64 `json:"income"`
}

type DailyIncome struct {
	Date   *gtime.Time `json:"date"`
	Income float64     `json:"income"`
}

type TimeRange struct {
	StartDate *gtime.Time `json:"startDate"`
	EndDate   *gtime.Time `json:"endDate"`
}
