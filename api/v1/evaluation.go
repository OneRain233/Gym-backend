package v1

import "github.com/gogf/gf/v2/frame/g"

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
