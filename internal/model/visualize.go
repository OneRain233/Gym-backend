package model

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
