package payment

import (
	"Gym-backend/internal/consts"
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

func (s *sPayment) CreatePayment(ctx context.Context, form *model.CreatePaymentForm) (response *model.ResponsePaymentForm, err error) {
	response = &model.ResponsePaymentForm{}
	order, err := service.Order().GetOrderByOrderCode(ctx, form.OrderCode)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}

	wallet, err := service.Wallet().GetWallet(ctx)
	if err != nil {
		return
	}
	if wallet == nil {
		err = gerror.New("wallet not found")
		return
	}
	paymentRecord := entity.Payment{
		WalletId:    wallet.Id,
		OrderId:     order.Id,
		PaymentCode: s.GeneratePaymentCode(),
		Time:        gtime.Now(),
		Amount:      order.Amount,
		PaymentType: form.PaymentType,
	}
	_, err = dao.Payment.Ctx(ctx).Insert(paymentRecord)
	if err != nil {
		return
	}

	// TODO: payment type and pay logic
	if paymentRecord.PaymentType == consts.PaymentTypeWallet {

	} else if paymentRecord.PaymentType == consts.PaymentTypeCard {
		if form.CardId == 0 {
			err = gerror.New("card id is required")
			return
		}
		cardPaymentForm := model.CardPayForm{
			CardId:  form.CardId,
			Amount:  order.Amount,
			OrderId: order.Id,
		}
		err = service.Card().Pay(ctx, &cardPaymentForm)
		if err != nil {
			return
		}

	}

	response.PaymentCode = paymentRecord.PaymentCode
	response.Amount = paymentRecord.Amount
	response.PaymentType = paymentRecord.PaymentType
	response.OrderCode = order.OrderCode
	response.Status = consts.PaymentSuccess

	return

}

func (s *sPayment) GeneratePaymentCode() string {
	// YearMonthDay + 8 digits
	return gtime.Now().Format("Ymd") + string(rand.Int())
}
