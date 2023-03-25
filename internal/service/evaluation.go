// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"context"
)

type (
	IEvaluation interface {
		AddEvaluation(ctx context.Context, form *model.AddEvaluationForm) error
		ValidateAddEvaluation(ctx context.Context, form *model.AddEvaluationForm) error
		ValidateUpdateEvaluation(ctx context.Context, form *model.UpdateEvaluationForm) error
		GetEvaluationByFacilityIdAndUserId(ctx context.Context, userId int, facilityId int) (evaluation *entity.Evaluation, err error)
		GetAllEvaluations(ctx context.Context) (evaluations []*entity.Evaluation, err error)
		GetEvaluationsByUserId(ctx context.Context, userId int) (evaluations []*entity.Evaluation, err error)
		GetEvaluationByFacilityId(ctx context.Context, facilityId int) (evaluations []*entity.Evaluation, err error)
		DeleteEvaluationByUserIdAndFacilityId(ctx context.Context, userId int, facilityId int) error
		DeleteEvaluationById(ctx context.Context, id int) error
		GetEvaluationById(ctx context.Context, id int) (evaluation *entity.Evaluation, err error)
		UpdateEvaluation(ctx context.Context, form *model.UpdateEvaluationForm) error
	}
)

var (
	localEvaluation IEvaluation
)

func Evaluation() IEvaluation {
	if localEvaluation == nil {
		panic("implement not found for interface IEvaluation, forgot register?")
	}
	return localEvaluation
}

func RegisterEvaluation(i IEvaluation) {
	localEvaluation = i
}
