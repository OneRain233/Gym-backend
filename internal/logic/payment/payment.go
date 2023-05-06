package payment

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"fmt"
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

func (s *sPayment) CreatePaymentForRegularOrder(ctx context.Context, form *model.CreatePaymentForm) (response *model.ResponsePaymentForm, err error) {
	response = &model.ResponsePaymentForm{}
	orders, err := service.Order().GetRegularOrdersByOrderCode(ctx, form.OrderCode)
	if err != nil {
		return
	}
	if orders == nil {
		err = gerror.New("order not found")
		return
	}
	amount := 0.0
	for _, order := range orders {
		amount += order.Amount
		if order.Status != consts.OrderStatusWaitingPayment {
			return nil, gerror.New("order status not valid")
		}
	}
	if form.PaymentType == consts.PaymentTypeWallet {
		wallet, err := service.Wallet().GetWallet(ctx)
		if err != nil {
			return nil, err
		}
		if wallet.Amount < amount {
			return nil, gerror.New("not enough money")
		}
	} else if form.PaymentType == consts.PaymentTypeCard {
		card, err := service.Card().GetCardById(ctx, form.CardId)
		if err != nil {
			return nil, err
		}
		if card == nil {
			return nil, gerror.New("card not found")
		}
		if card.Amount < amount {
			return nil, gerror.New("not enough money")
		}
	} else {
		return nil, gerror.New("payment type not valid")
	}

	for _, order := range orders {
		newForm := &model.CreatePaymentForm{
			OrderCode:   order.OrderCode,
			PaymentType: form.PaymentType,
			CardId:      form.CardId,
		}
		_, err1 := s.CreatePayment(ctx, newForm)
		if err1 != nil {
			return nil, err1
		}
		//response = resp
	}
	return
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

	if order.Status != consts.OrderStatusWaitingPayment {
		err = gerror.New("order status not valid")
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
	if paymentRecord.PaymentType != form.PaymentType {
		// update payment type
		paymentRecord.PaymentType = form.PaymentType
		_, err = dao.Payment.Ctx(ctx).Where("id", paymentRecord.Id).Update(paymentRecord)
		if err != nil {
			err = gerror.New("payment update failed")
			return
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
	fmt.Println(paymentRecord.PaymentType)
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
	} else if paymentRecord.PaymentType == consts.PaymentTypeCash {
		// do nothing
	} else {
		err = gerror.New("payment type not valid")
		return
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

func (s *sPayment) CreatePaymentForRefund(ctx context.Context, form *model.RefundPaymentForm) error {
	paymentEntity := &entity.Payment{
		WalletId:    form.OriginalPayment.WalletId,
		OrderId:     form.Order.Id,
		PaymentCode: form.OriginalPayment.PaymentCode,
		Time:        gtime.Now(),
		Amount:      form.OriginalPayment.Amount,
		PaymentType: consts.PaymentRefund,
		Status:      consts.PaymentSuccess,
	}
	_, err := dao.Payment.Ctx(ctx).Insert(paymentEntity)
	if err != nil {
		return err
	}
	return nil
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
	payment, err := s.GetPaymentById(ctx, paymentId)
	if err != nil {
		return
	}
	payment.Status = status
	_, err = dao.Payment.Ctx(ctx).Where("id", paymentId).Update(payment)
	if err != nil {
		return
	}
	return
}

func (s *sPayment) UpdatePaymentStatusByOrderId(ctx context.Context, orderId int, status int) error {
	payment, err := s.GetPaymentByOrderId(ctx, orderId)
	if err != nil {
		return err
	}
	payment.Status = status
	_, err = dao.Payment.Ctx(ctx).Where("order_id", orderId).Update(payment)
	if err != nil {
		return err
	}
	return nil
}
