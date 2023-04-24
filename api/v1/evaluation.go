package v1

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddEvaluationReq struct {
	g.Meta      `path:"/evaluation/add" tags:"Evaluation" method:"post" summary:"Add Evaluation"`
	FacilityId  int      `json:"facility_id" v:"required#facility_id is required"`
	Score       int      `json:"score" v:"required#score is required"`
	IsAnonymous int      `json:"is_anonymous" v:"required#is_anonymous is required"`
	Description string   `json:"description" v:"required#description is required"`
	Images      []string `json:"images"`
	Videos      []string `json:"videos"`
}

type AddEvaluationRes struct {
}

type DeleteOwnEvaluationReq struct {
	g.Meta `path:"/evaluation/delete" tags:"Evaluation" method:"post" summary:"Delete Own Evaluation"`
	Id     int `json:"id" v:"required#id is required"`
}

type DeleteOwnEvaluationRes struct {
}

type UpdateEvaluationReq struct {
	g.Meta      `path:"/evaluation/update" tags:"Evaluation" method:"post" summary:"Update Evaluation"`
	Id          int      `json:"id" v:"required#id is required"`
	FacilityId  int      `json:"facility_id" v:"required#facility_id is required"`
	Score       int      `json:"score" v:"required#score is required"`
	IsAnonymous int      `json:"is_anonymous" v:"required#is_anonymous is required"`
	Description string   `json:"description" v:"required#description is required"`
	Images      []string `json:"images"`
	Videos      []string `json:"videos"`
}

type UpdateEvaluationRes struct {
}

type GetEvaluationsReq struct {
	g.Meta `path:"/evaluation/get" tags:"Evaluation" method:"post" summary:"Get Evaluations"`
}

type GetEvaluationsRes struct {
	Evaluations []*entity.Evaluation `json:"evaluations"`
}

type AdminGetEvaluationReq struct {
	g.Meta     `path:"/evaluation/admin/get" tags:"Evaluation" method:"post" summary:"Admin Get Evaluations"`
	UserId     int `json:"user_id"`
	FacilityId int `json:"facility_id"`
}

type AdminGetEvaluationRes struct {
	Evaluations []*entity.Evaluation `json:"evaluations"`
}

type UserGetEvaluationsReq struct {
	g.Meta     `path:"/evaluation/user/get" tags:"Evaluation" method:"post" summary:"User Get Evaluations"`
	FacilityId int `json:"facility_id"`
}

type UserGetEvaluationRes struct {
	Evaluations []*model.SafeEvaluation `json:"evaluations"`
}
