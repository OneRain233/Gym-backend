package payment

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"math/rand"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sPayment struct{}

func init() {
	service.RegisterPayment(New())
}

func New() *sPayment {
	return &sPayment{}
}

func (s *sPayment) CreatePayment(ctx context.Context, form *model.CreatePaymentForm) error {
	if form.WalletId == 0 {
		return gerror.New("wallet id is required")
	}
	if form.Amount == 0 {
		return gerror.New("amount is required")
	}
	if form.Amount < 0 {
		return gerror.New("amount must be positive")
	}

	payment := entity.Payment{
		WalletId:    form.WalletId,
		OrderId:     form.OrderId,
		PaymentCode: string(rand.Int()),
		Time:        gtime.Now(),
		Amount:      form.Amount,
		PaymentType: form.PaymentType,
	}
	_, err := dao.Payment.Ctx(ctx).Insert(&payment)
	return err

}
