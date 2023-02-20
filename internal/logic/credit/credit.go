package credit

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sCredit struct{}

func init() {
	service.RegisterCredit(New())
}

func New() *sCredit {
	return &sCredit{}
}

func (s *sCredit) GetCredit(ctx context.Context) (credit *entity.Credit, err error) {
	userId := service.Session().GetUser(ctx).Id
	err = dao.Credit.Ctx(ctx).Where(dao.Credit.Columns().UserId, userId).Scan(&credit)
	if err != nil {
		return
	}
	return
}

func (s *sCredit) UpdateAmount(ctx context.Context, amount int) error {
	userId := service.Session().GetUser(ctx).Id
	var credit *entity.Credit
	err := dao.Credit.Ctx(ctx).Where("userId", userId).Scan(&credit)
	if err != nil {
		return err
	}
	credit.Credit = amount
	_, err = dao.Credit.Ctx(ctx).Where("userId", userId).Update(&credit)
	return err
}

func (s *sCredit) GetAllCredit(ctx context.Context) (credits []*entity.Credit, err error) {
	err = dao.Credit.Ctx(ctx).Scan(&credits)
	if err != nil {
		return
	}
	return
}

func (s *sCredit) GetCreditByUserId(ctx context.Context, userId int) (credit *entity.Credit, err error) {
	err = dao.Credit.Ctx(ctx).Where(dao.Credit.Columns().UserId, userId).Scan(&credit)
	if err != nil {
		return
	}
	return
}

func (s *sCredit) UpdateCreditByUserId(ctx context.Context, userId int, amount int) error {
	var credit *entity.Credit
	err := dao.Credit.Ctx(ctx).Where(dao.Credit.Columns().UserId, userId).Scan(&credit)
	if err != nil {
		return err
	}
	credit.Credit = amount
	_, err = dao.Credit.Ctx(ctx).Where(dao.Credit.Columns().UserId, userId).Update(&credit)
	return err
}

func (s *sCredit) CreateCreditForUser(ctx context.Context, userId int) error {
	credit := &entity.Credit{
		UserId: userId,
		Credit: 0,
	}
	_, err := dao.Credit.Ctx(ctx).Insert(credit)
	return err
}
