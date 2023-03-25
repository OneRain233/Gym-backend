package evaluation

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sEvaluation struct{}

func init() {
	service.RegisterEvaluation(New())
}

func New() *sEvaluation {
	return &sEvaluation{}
}

func (c *sEvaluation) AddEvaluation(ctx context.Context, form *model.AddEvaluationForm) error {
	if err := c.ValidateEvaluation(ctx, form); err != nil {
		return err
	}
	userId := form.UserId
	facilityId := form.FacilityId
	evaluation := &entity.Evaluation{
		UserId:      userId,
		FacilityId:  facilityId,
		Score:       form.Score,
		Description: form.Description,
		Anonymous:   form.IsAnonymous,
		Images:      form.Images,
		Videos:      form.Videos,
	}
	_, err := dao.Evaluation.Ctx(ctx).Data(evaluation).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (c *sEvaluation) ValidateEvaluation(ctx context.Context, form *model.AddEvaluationForm) error {
	evaluations, err := c.GetEvaluationByFacilityIdAndUserId(ctx, form.UserId, form.FacilityId)
	if err != nil {
		return err
	}
	if evaluations != nil {
		return gerror.New("You have already evaluated this facility")
	}
	facility, err := service.Facility().GetFacilityById(ctx, form.FacilityId)
	if err != nil {
		return err
	}
	if facility.Facility == nil {
		return gerror.New("Facility does not exist")
	}
	if form.IsAnonymous != 0 && form.IsAnonymous != 1 {
		return gerror.New("IsAnonymous should be 0 or 1")
	}
	if form.Score < 0 || form.Score > 5 {
		return gerror.New("Score should be between 0 and 5")
	}
	return nil
}

func (c *sEvaluation) GetEvaluationByFacilityIdAndUserId(ctx context.Context, userId int, facilityId int) (evaluation *entity.Evaluation, err error) {
	err = dao.Evaluation.Ctx(ctx).Where("user_id", userId).Where("facility_id", facilityId).Scan(&evaluation)
	if err != nil {
		return
	}
	return
}
