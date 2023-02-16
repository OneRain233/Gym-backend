package payment

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"strconv"

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

	// check if there is record in db
	var paymentRecord *entity.Payment
	paymentRecord, err = s.GetPaymentByOrderId(ctx, order.Id)
	if err != nil {
		paymentRecord = &entity.Payment{
			WalletId:    wallet.Id,
			OrderId:     order.Id,
			PaymentCode: s.GeneratePaymentCode(),
			Time:        gtime.Now(),
			Amount:      order.Amount,
			PaymentType: form.PaymentType,
			Status:      consts.PaymentPending,
		}
		res, err1 := dao.Payment.Ctx(ctx).Insert(paymentRecord)
		paymentId, err2 := res.LastInsertId()
		paymentRecord.Id = int(paymentId)
		if err1 != nil || err2 != nil {
			err = gerror.New("payment create failed")
		}
	}

	// if this payment is already paid
	if paymentRecord.Status == consts.PaymentSuccess {
		response.PaymentCode = paymentRecord.PaymentCode
		response.Amount = paymentRecord.Amount
		response.PaymentType = paymentRecord.PaymentType
		response.OrderCode = order.OrderCode
		response.Status = consts.PaymentSuccess
		return
	}

	// TODO: payment type and pay logic
	if paymentRecord.PaymentType == consts.PaymentTypeWallet {
		walletPayForm := model.WalletPayForm{
			OrderId: order.Id,
			Amount:  order.Amount,
		}
		err = service.Wallet().Pay(ctx, &walletPayForm)
		if err != nil {
			// update payment status
			paymentRecord.Status = consts.PaymentFail
			_, _ = dao.Payment.Ctx(ctx).Where("id", paymentRecord.Id).Update(paymentRecord)
			return
		}
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
			// update payment status
			paymentRecord.Status = consts.PaymentFail
			_, _ = dao.Payment.Ctx(ctx).Where("id", paymentRecord.Id).Update(paymentRecord)
			return
		}
	}

	// update payment status
	paymentRecord.Status = consts.PaymentSuccess
	_, err = dao.Payment.Ctx(ctx).Where("id", paymentRecord.Id).Update(paymentRecord)
	if err != nil {
		return
	}
	order.Status = consts.OrderStatusPending
	_, err = dao.Order.Ctx(ctx).Where("id", order.Id).Update(order)
	if err != nil {
		return
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
	return gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
}

// GetPaymentByPaymentCode get payment by payment code
func (s *sPayment) GetPaymentByPaymentCode(ctx context.Context, paymentCode string) (payment *entity.Payment, err error) {
	payment = &entity.Payment{}
	err = dao.Payment.Ctx(ctx).Where("payment_code", paymentCode).Scan(payment)
	if err != nil {
		return
	}
	return
}

func (s *sPayment) GetPaymentById(ctx context.Context, paymentId int) (payment *entity.Payment, err error) {
	payment = &entity.Payment{}
	err = dao.Payment.Ctx(ctx).Where("id", paymentId).Scan(payment)
	if err != nil {
		return
	}
	return
}

func (s *sPayment) GetPaymentByOrderId(ctx context.Context, orderId int) (payment *entity.Payment, err error) {
	payment = &entity.Payment{}
	err = dao.Payment.Ctx(ctx).Where("order_id", orderId).OmitEmpty().Scan(payment)
	if err != nil {
		return
	}
	return
}

func (s *sPayment) GetPaymentByUserId(ctx context.Context, userId int) (payment []*entity.Payment, err error) {
	// check if user exists
	userEntity, err := service.User().GetUserByID(ctx, uint(userId))
	if err != nil || userEntity == nil {
		err = gerror.New("user not found")
		return
	}
	payment = []*entity.Payment{}
	wallet := &entity.Wallet{}
	err = dao.Wallet.Ctx(ctx).Where("user_id", userEntity.Id).Scan(&wallet)
	if wallet == nil || err != nil {
		err = gerror.New("wallet not found")
		return
	}
	err = dao.Payment.Ctx(ctx).Where("wallet_id", wallet.Id).Scan(&payment)
	if payment == nil {
		err = gerror.New("payment not found")
		return
	}

	return
}

func (s *sPayment) GetAllPayment(ctx context.Context) (payments []*entity.Payment, err error) {
	payments = []*entity.Payment{}
	err = dao.Payment.Ctx(ctx).Scan(&payments)
	if err != nil {
		return
	}
	return
}

func (s *sPayment) GetPaymentByOrderCode(ctx context.Context, orderCode string) (payment *entity.Payment, err error) {
	var order *entity.Order
	order, err = service.Order().GetOrderByOrderCode(ctx, orderCode)
	if order == nil {
		err = gerror.New("order not found")
		return
	}
	if err != nil {
		return
	}
	orderId := order.Id
	err = dao.Payment.Ctx(ctx).Where("order_id", orderId).Scan(&payment)
	if payment == nil {
		err = gerror.New("payment not found")
		return
	}
	if err != nil {
		return
	}
	return
}

func (s *sPayment) UpdatePaymentStatus(ctx context.Context, paymentId int, status int) (err error) {
	payment := &entity.Payment{
		Status: status,
	}
	_, err = dao.Payment.Ctx(ctx).Where("id", paymentId).Update(payment)
	if err != nil {
		return
	}
	return
}
