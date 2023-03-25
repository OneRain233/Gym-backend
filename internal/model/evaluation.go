package model

type AddEvaluationForm struct {
	UserId      int    `json:"user_id"`
	FacilityId  int    `json:"facility_id"`
	Score       int    `json:"score"`
	Description string `json:"description"`
	IsAnonymous int    `json:"is_anonymous"`
	Images      string `json:"images"`
	Videos      string `json:"videos"`
}
