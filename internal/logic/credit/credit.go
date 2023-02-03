package credit

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sCredit struct{}

func New() *sCredit {
	return &sCredit{}
}

func (s *sCredit) GetCredit(ctx context.Context) (credit *entity.Credit, err error) {
	userId := service.Session().GetUser(ctx).Id
	err = dao.Credit.Ctx(ctx).Where("userId", userId).Scan(&credit)
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
