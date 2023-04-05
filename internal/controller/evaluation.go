package controller

import (
	v1 "Gym-backend/api/v1"
	"Gym-backend/internal/model"
	"Gym-backend/internal/service"
	"context"
	"strings"
)

var Evaluation = cEvaluation{}
var EvaluationAdmin = cEvaluationAdmin{}

type cEvaluation struct{}
type cEvaluationAdmin struct{}

func (c *cEvaluation) GetEvaluations(ctx context.Context, req *v1.GetEvaluationsReq) (res *v1.GetEvaluationsRes, err error) {
	res = &v1.GetEvaluationsRes{}
	user := service.Session().GetUser(ctx)
	userId := user.Id
	res.Evaluations, err = service.Evaluation().GetEvaluationsByUserId(ctx, userId)
	if err != nil {
		return
	}
	return
}

func (s *cEvaluationAdmin) GetEvaluations(ctx context.Context, req *v1.AdminGetEvaluationReq) (res *v1.AdminGetEvaluationRes, err error) {
	res = &v1.AdminGetEvaluationRes{}
	userId := req.UserId
	facilityId := req.FacilityId
	if userId == 0 && facilityId == 0 {
		evaluation, err1 := service.Evaluation().GetAllEvaluations(ctx)
		if evaluation != nil {
			res.Evaluations = evaluation
		}
		err = err1
	} else if facilityId == 0 {
		evaluations, err1 := service.Evaluation().GetEvaluationsByUserId(ctx, userId)
		res.Evaluations = evaluations
		err = err1
	} else if userId == 0 {
		evaluations, err1 := service.Evaluation().GetEvaluationByFacilityId(ctx, facilityId)
		res.Evaluations = evaluations
		err = err1
	}
	if err != nil {
		return
	}
	return
}

func (c *cEvaluation) AddEvaluation(ctx context.Context, req *v1.AddEvaluationReq) (res *v1.AddEvaluationRes, err error) {
	res = &v1.AddEvaluationRes{}
	videos := ""
	videos = strings.Join(req.Videos, ",")

	images := ""
	images = strings.Join(req.Images, ",")

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

func (c *cEvaluation) DeleteOwnEvaluation(ctx context.Context, req *v1.DeleteOwnEvaluationReq) (res *v1.DeleteOwnEvaluationRes, err error) {
	res = &v1.DeleteOwnEvaluationRes{}
	err = service.Evaluation().DeleteEvaluationById(ctx, req.Id)
	if err != nil {
		return
	}
	return
}

func (c *cEvaluation) UpdateEvaluation(ctx context.Context, req *v1.UpdateEvaluationReq) (res *v1.UpdateEvaluationRes, err error) {
	res = &v1.UpdateEvaluationRes{}
	videos := ""
	videos = strings.Join(req.Videos, ",")

	images := ""
	images = strings.Join(req.Images, ",")

	form := &model.UpdateEvaluationForm{
		Id:          req.Id,
		Score:       req.Score,
		Description: req.Description,
		IsAnonymous: req.IsAnonymous,
		Images:      images,
		Videos:      videos,
	}
	err = service.Evaluation().UpdateEvaluation(ctx, form)
	if err != nil {
		return
	}
	return
}
