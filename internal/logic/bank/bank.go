package bank

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
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
