package bank

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sBank struct {
}

func init() {
	service.RegisterBank(New())
}

func New() *sBank {
	return &sBank{}
}

func (s *sBank) GetBanks(ctx context.Context) (banks []*entity.Bank, err error) {
	err = dao.Bank.Ctx(ctx).Scan(&banks)
	if err != nil {
		return
	}
	return
}

func (s *sBank) GetBankById(ctx context.Context, bankId int) (bank *entity.Bank, err error) {
	bank = &entity.Bank{}
	err = dao.Bank.Ctx(ctx).Where("id", bankId).Scan(bank)
	if err != nil {
		return
	}
	return
}

func (s *sBank) AddBank(ctx context.Context, form *model.AddBankForm) error {
	if form.Name == "" {
		return gerror.New("name is empty")
	}
	bank := entity.Bank{
		Name: form.Name,
	}
	_, err := dao.Bank.Ctx(ctx).Insert(&bank)
	return err
}

func (s *sBank) UpdateBank(ctx context.Context, form *model.UpdateBankForm) error {
	if form.Id == 0 {
		return gerror.New("id is 0")
	}
	if form.Name == "" {
		return gerror.New("name is empty")
	}
	var bank *entity.Bank
	err := dao.Bank.Ctx(ctx).Where("id", form.Id).Scan(&bank)
	if err != nil {
		return err
	}
	if bank == nil {
		return gerror.New("bank not found")
	}
	bank.Name = form.Name
	_, err = dao.Bank.Ctx(ctx).Where("id", form.Id).Update(bank)
	return err
}
