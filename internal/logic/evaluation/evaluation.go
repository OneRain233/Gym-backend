package evaluation

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gtime"

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
	if err := c.ValidateAddEvaluation(ctx, form); err != nil {
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
		Time:        gtime.Now(),
	}
	_, err := dao.Evaluation.Ctx(ctx).Data(evaluation).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (c *sEvaluation) ValidateAddEvaluation(ctx context.Context, form *model.AddEvaluationForm) error {
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

func (c *sEvaluation) ValidateUpdateEvaluation(ctx context.Context, form *model.UpdateEvaluationForm) error {
	evaluation, err := c.GetEvaluationById(ctx, form.Id)
	if err != nil {
		return err
	}
	if evaluation == nil {
		return gerror.New("Evaluation does not exist")
	}
	if evaluation.UserId != service.Session().GetUser(ctx).Id && service.Session().GetUser(ctx).Role != consts.RoleAdmin {
		return gerror.New("You are not allowed to update this evaluation")
	}
	facility, err := service.Facility().GetFacilityById(ctx, evaluation.FacilityId)
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

func (c *sEvaluation) GetAllEvaluations(ctx context.Context) (evaluations []*entity.Evaluation, err error) {
	err = dao.Evaluation.Ctx(ctx).Scan(&evaluations)
	if err != nil {
		return
	}
	return
}

func (c *sEvaluation) GetEvaluationsByUserId(ctx context.Context, userId int) (evaluations []*entity.Evaluation, err error) {
	err = dao.Evaluation.Ctx(ctx).Where("user_id", userId).Scan(&evaluations)
	if err != nil {
		return
	}
	return
}

func (c *sEvaluation) GetEvaluationByFacilityId(ctx context.Context, facilityId int) (evaluations []*entity.Evaluation, err error) {
	err = dao.Evaluation.Ctx(ctx).Where("facility_id", facilityId).Scan(&evaluations)
	if err != nil {
		return
	}
	return
}

func (c *sEvaluation) DeleteEvaluationByUserIdAndFacilityId(ctx context.Context, userId int, facilityId int) error {
	_, err := dao.Evaluation.Ctx(ctx).Where("user_id", userId).Where("facility_id", facilityId).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (c *sEvaluation) DeleteEvaluationById(ctx context.Context, id int) error {
	// search if the evaluation is belong to the user
	user := service.Session().GetUser(ctx)
	userId := user.Id
	if userId == 0 {
		return gerror.New("Please login first")
	}
	evaluation, err := c.GetEvaluationById(ctx, id)
	if err != nil {
		return err
	}
	if evaluation == nil {
		return gerror.New("Evaluation does not exist")
	}
	if evaluation.UserId != userId && user.Role != consts.RoleAdmin {
		return gerror.New("You are not allowed to delete this evaluation")
	}

	_, err = dao.Evaluation.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (c *sEvaluation) GetEvaluationById(ctx context.Context, id int) (evaluation *entity.Evaluation, err error) {
	err = dao.Evaluation.Ctx(ctx).Where("id", id).Scan(&evaluation)
	if err != nil {
		return
	}
	return
}

func (c *sEvaluation) UpdateEvaluation(ctx context.Context, form *model.UpdateEvaluationForm) error {
	if err := c.ValidateUpdateEvaluation(ctx, form); err != nil {
		return err
	}
	evaluationId := form.Id
	var evaluation *entity.Evaluation
	evaluation, err := c.GetEvaluationById(ctx, evaluationId)
	if err != nil {
		return err
	}
	if evaluation == nil {
		return gerror.New("Evaluation does not exist")
	}

	// difference update
	evaluation.Score = form.Score
	evaluation.Anonymous = form.IsAnonymous
	evaluation.Description = form.Description
	evaluation.Images = form.Images
	evaluation.Videos = form.Videos
	evaluation.Time = gtime.Now()

	_, err = dao.Evaluation.Ctx(ctx).Data(evaluation).Where("id", evaluationId).Update()
	if err != nil {
		return err
	}

	return nil

}

func (c *sEvaluation) GetFacilityScore(ctx context.Context, facilityId int) (score int, err error) {
	evaluations, err := c.GetEvaluationByFacilityId(ctx, facilityId)
	if err != nil {
		return
	}
	if len(evaluations) == 0 {
		return 5, nil
	}
	var total int
	for _, evaluation := range evaluations {
		total += evaluation.Score
	}
	score = total / len(evaluations)
	return
}

func (c *sEvaluation) FetchUsername(ctx context.Context, evaluation *entity.Evaluation) (user *entity.User, err error) {
	user, err = service.User().GetUserById(ctx, uint(evaluation.UserId))
	if err != nil {
		return
	}
	if user == nil {
		return nil, gerror.New("User does not exist")
	}
	return user, nil
}
