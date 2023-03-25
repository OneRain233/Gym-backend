package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
)

var Evaluation = cEvaluation{}
var EvaluationAdmin = cEvaluationAdmin{}

type cEvaluation struct{}
type cEvaluationAdmin struct{}

func (c *cEvaluation) AddEvaluation(ctx context.Context, req *v1.AddEvaluationReq) (res *v1.AddEvaluationRes, err error) {
	res = &v1.AddEvaluationRes{}
	videos := ""
	if len(req.Videos) != 0 {
		for _, v := range req.Videos {
			videos += v + ","
		}
	}

	images := ""
	if len(req.Images) != 0 {
		for _, v := range req.Images {
			images += v + ","
		}
	}
	form := &model.AddEvaluationForm{
		UserId:      service.Session().GetUser(ctx).Id,
		FacilityId:  req.FacilityId,
		Score:       req.Score,
		Description: req.Description,
		IsAnonymous: req.IsAnonymous,
		Images:      images,
		Videos:      videos,
	}
	err = service.Evaluation().AddEvaluation(ctx, form)
	if err != nil {
		return
	}
	return
}
