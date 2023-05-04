package model

import "Gym-backend/internal/model/entity"

type AddEvaluationForm struct {
	UserId      int    `json:"user_id"`
	FacilityId  int    `json:"facility_id"`
	Score       int    `json:"score"`
	Description string `json:"description"`
	IsAnonymous int    `json:"is_anonymous"`
	Images      string `json:"images"`
	Videos      string `json:"videos"`
}

type UpdateEvaluationForm struct {
	Id          int    `json:"id"`
	FacilityId  int    `json:"facility_id"`
	Score       int    `json:"score"`
	Description string `json:"description"`
	IsAnonymous int    `json:"is_anonymous"`
	Images      string `json:"images"`
	Videos      string `json:"videos"`
}

type SafeEvaluation struct {
	Evaluation *entity.Evaluation `json:"evaluation"`
	UserName   string             `json:"username"`
	Avatar     string             `json:"avatar"`
}
