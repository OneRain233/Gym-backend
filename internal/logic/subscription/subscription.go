package subscription

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSubscription struct {
}

func init() {
	service.RegisterSubscription(New())
}

func New() *sSubscription {
	return &sSubscription{}
}

func (s *sSubscription) GetSubscriptionEndDayByUserId(ctx context.Context, userId int) (res *gtime.Time, err error) {
	res = gtime.Now()
	var allSubscription []*entity.Subscription
	err = dao.Subscription.Ctx(ctx).Where("user_id", userId).Scan(&allSubscription)
	if err != nil {
		return
	}
	for _, subscription := range allSubscription {
		endTime := subscription.EndTime
		if endTime.Timestamp() > res.Timestamp() {
			res = endTime
		}
	}
	return
}

func (s *sSubscription) GetSubscriptionListByUserId(ctx context.Context, userId int) (res []*entity.Subscription, err error) {
	res = make([]*entity.Subscription, 0)
	err = dao.Subscription.Ctx(ctx).Where("user_id", userId).Scan(&res)
	return
}

func (s *sSubscription) CreateSubscription(ctx context.Context, form *model.CreateSubscriptionForm) error {
	userId := form.UserId
	paymentType := form.PaymentType
	endDay, err := s.GetSubscriptionEndDayByUserId(ctx, userId)
	if err != nil {
		return err
	}
	// not expired
	if endDay.Timestamp() > gtime.Now().Timestamp() {
		return gerror.New("You have not expired subscription")
	} else {
		// create order
		//orderCode := s.GenerateOrderCode()
		if paymentType == consts.PaymentTypeCard {
			cardPaymentForm := &model.CardPayForm{
				CardId: form.CardId,
				Amount: 1, // TODO
			}
			err = service.Card().PayForSubscription(ctx, cardPaymentForm)
			if err != nil {
				return err
			}
		} else if paymentType == consts.PaymentTypeWallet {
			walletPayForm := &model.WalletPayForm{
				Amount: 1, // TODO
			}
			err = service.Wallet().PayForSubscription(ctx, walletPayForm)
			if err != nil {
				return err
			}
		} else {
			return gerror.New("Payment type not found")
		}
		_, err = dao.Subscription.Ctx(ctx).Data(&entity.Subscription{
			UserId:    userId,
			StartTime: gtime.Now(),
			EndTime:   gtime.Now().AddDate(0, 0, form.Days),
		}).Insert()
		return err
	}

}

func (s *sSubscription) GenerateOrderCode() string {
	// YearMonthDay + 8 digits
	return gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
}
