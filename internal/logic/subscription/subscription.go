package subscription

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"

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

func (c *sSubscription) InitSubscriptionSTypesToCache(ctx context.Context) error {
	err := service.Cache().Delete(ctx, "subscription_type")
	if err != nil {
		return err
	}

	var subscriptionTypes []*entity.SubscriptionType
	err = dao.SubscriptionType.Ctx(ctx).Scan(&subscriptionTypes)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "Get Subscriptions", subscriptionTypes)

	for _, subscriptionType := range subscriptionTypes {
		g.Log().Info(ctx, "Get Subscription", subscriptionType)
		err = service.Cache().Push(ctx, "subscription_type", subscriptionType)
		if err != nil {
			continue
		}
		err = service.Cache().Set(ctx, "subscription_type_"+strconv.Itoa(subscriptionType.Id), subscriptionType, 0)
		if err != nil {
			continue
		}
	}
	return nil
}

func (s *sSubscription) GetTypesOfSubscription(ctx context.Context) (res []*entity.SubscriptionType, err error) {
	//res = make([]*entity.SubscriptionType, 0)
	var cacheRes interface{}
	cacheRes, err = service.Cache().GetList(ctx, "subscription_type")
	g.Log().Info(ctx, cacheRes)
	if cacheRes != nil && len(cacheRes.([]interface{})) > 0 {
		g.Log().Debugf(ctx, "cache exists", len(cacheRes.([]interface{})))
		for _, v := range cacheRes.([]interface{}) {
			var subscriptionType *entity.SubscriptionType
			err = gconv.Struct(v, &subscriptionType)
			if err != nil {
				return
			}
			res = append(res, subscriptionType)
		}
		return
	}
	res = make([]*entity.SubscriptionType, 0)
	err = dao.SubscriptionType.Ctx(ctx).Scan(&res)
	_ = s.InitSubscriptionSTypesToCache(ctx)
	return
}

func (s *sSubscription) GetSubscriptionPermissionById(ctx context.Context, permissionId int) (res *entity.SubscriptionPermission, err error) {
	res = &entity.SubscriptionPermission{}
	err = dao.SubscriptionPermission.Ctx(ctx).Where("id", permissionId).Scan(res)
	return
}

func (s *sSubscription) GetSubscriptionTypeById(ctx context.Context, id int) (res *entity.SubscriptionType, err error) {
	res = &entity.SubscriptionType{}
	// check cache
	var cacheRes interface{}
	cacheRes, err = service.Cache().Get(ctx, "subscription_type_"+strconv.Itoa(id))
	if cacheRes != nil {
		err = gconv.Struct(cacheRes, &res)
		if err != nil {
			return
		}
		return
	}

	err = dao.SubscriptionType.Ctx(ctx).Where("id", id).Scan(res)
	return
}

func (s *sSubscription) GetSubscriptionEndDayByUserId(ctx context.Context, userId int) (sub *entity.Subscription, res *gtime.Time, err error) {
	res = gtime.NewFromStr("1970-01-01 00:00:00")
	var allSubscription []*entity.Subscription
	err = dao.Subscription.Ctx(ctx).Where("user_id", userId).Scan(&allSubscription)
	if err != nil {
		return
	}
	for _, subscription := range allSubscription {
		endTime := subscription.EndTime
		if endTime.Timestamp() > res.Timestamp() && subscription.Status == consts.SubscriptionStatusNormal {
			res = endTime
			sub = subscription
		}
	}
	if res.Timestamp() == gtime.NewFromStr("1970-01-01 00:00:00").Timestamp() {
		return nil, gtime.NewFromStr("1970-01-01 00:00:00"), nil
	}
	return
}

func (s *sSubscription) GetUnFinishedSubscription(ctx context.Context, userId int) (res *entity.Subscription, err error) {
	res = &entity.Subscription{}
	now := gtime.Now().Format("Y-m-d H:i:s")
	err = dao.Subscription.Ctx(ctx).Where("user_id", userId).Where("end_time > ?", now).Scan(res)
	return
}
func (s *sSubscription) GetSubscriptionListByUserId(ctx context.Context, userId int) (res []*entity.Subscription, err error) {
	res = make([]*entity.Subscription, 0)
	err = dao.Subscription.Ctx(ctx).Where("user_id", userId).Scan(&res)
	return
}

func (s *sSubscription) GetSubscriptionByUserId(ctx context.Context, userId int) (res *entity.Subscription, err error) {
	res = &entity.Subscription{}
	today := gtime.Now().Format("Y-m-d H:i:s")
	// find subscription that end time is greater than today
	err = dao.Subscription.Ctx(ctx).Where("user_id", userId).Where("end_time > ?", today).Scan(res)
	if res == nil {
		return nil, gerror.New("You have not subscription")
	}
	if err != nil {
		return nil, gerror.New("You have not subscription")
	}
	return
}

func (s *sSubscription) CreateSubscription(ctx context.Context, form *model.CreateSubscriptionForm) error {
	userId := form.UserId
	paymentType := form.PaymentType
	subscriptionType, err := s.GetSubscriptionTypeById(ctx, form.Type)
	if err != nil {
		return err
	}
	days := subscriptionType.Days
	amount := subscriptionType.Amount
	//subscription, err := s.GetUnFinishedSubscription(ctx, userId)
	sub, endDay, err := s.GetSubscriptionEndDayByUserId(ctx, userId)
	if err != nil {
		return err
	}
	// not expired
	fmt.Println(endDay)
	fmt.Println((endDay.Timestamp() > gtime.Now().Timestamp()))
	if endDay.Timestamp() > gtime.Now().Timestamp() && sub.Status == consts.SubscriptionStatusNormal {
		return gerror.New("You have not expired subscription")
	} else {
		// create order
		//orderCode := s.GenerateOrderCode()
		if paymentType == consts.PaymentTypeCard {
			cardPaymentForm := &model.CardPayForm{
				CardId: form.CardId,
				Amount: amount,
			}
			err = service.Card().PayForSubscription(ctx, cardPaymentForm)
			if err != nil {
				return err
			}
		} else if paymentType == consts.PaymentTypeWallet {
			walletPayForm := &model.WalletPayForm{
				Amount: amount,
			}
			err = service.Wallet().PayForSubscription(ctx, walletPayForm)
			if err != nil {
				return err
			}
		} else {
			return gerror.New("Payment type not found")
		}
		_, err = dao.Subscription.Ctx(ctx).Data(&entity.Subscription{
			UserId:           userId,
			StartTime:        gtime.Now(),
			EndTime:          gtime.Now().AddDate(0, 0, days),
			SubscriptionType: form.Type,
		}).Insert()
		return err
	}

}

func (s *sSubscription) UpdateSubscriptionStatus(ctx context.Context, subscriptionId int, status int) error {
	var subscription *entity.Subscription
	err := dao.Subscription.Ctx(ctx).Where("id", subscriptionId).Scan(&subscription)
	if err != nil {
		return err
	}
	if subscription == nil {
		return gerror.New("Subscription not found")
	}
	_, err = dao.Subscription.Ctx(ctx).Data(g.Map{
		"status": status,
	}).Where("id", subscriptionId).Update()

	return err
}

func (s *sSubscription) GenerateOrderCode() string {
	// YearMonthDay + 8 digits
	return gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
}

func (s *sSubscription) CancelSubscription(ctx context.Context) error {
	userId := service.Session().GetUser(ctx).Id
	subscription, _, err := s.GetSubscriptionEndDayByUserId(ctx, userId)
	if err != nil {
		return err
	}
	if subscription == nil {
		return gerror.New("You have not subscription")
	}
	if subscription.Status != consts.SubscriptionStatusNormal {
		return gerror.New("You have not subscription")
	}
	// update subscription status
	err = s.UpdateSubscriptionStatus(ctx, subscription.Id, consts.SubscriptionStatusCancel)
	if err != nil {
		return err
	}
	return nil
}

func (s *sSubscription) CheckExpiredSubscription(ctx context.Context) error {
	var subscriptions []*entity.Subscription
	err := dao.Subscription.Ctx(ctx).Where("end_time < ?", gtime.Now().Format("Y-m-d H:i:s")).Scan(&subscriptions)
	if err != nil {
		return err
	}
	for _, subscription := range subscriptions {
		err = s.UpdateSubscriptionStatus(ctx, subscription.Id, consts.SubscriptionStatusExpire)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sSubscription) CancelSubscriptionByUserId(ctx context.Context, userId int) error {
	subscription, endDay, err := s.GetSubscriptionEndDayByUserId(ctx, service.Session().GetUser(ctx).Id)
	if err != nil {
		return err
	}
	if endDay.Timestamp() > gtime.Now().Timestamp() && subscription.Status == consts.SubscriptionStatusNormal {
		err = s.UpdateSubscriptionStatus(ctx, subscription.Id, consts.SubscriptionStatusCancel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sSubscription) GetAllSubscriptions(ctx context.Context) ([]*entity.Subscription, error) {
	var subscriptions []*entity.Subscription
	err := dao.Subscription.Ctx(ctx).Scan(&subscriptions)
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (s *sSubscription) UpdateSubscriptionType(ctx context.Context, form *model.UpdateSubscriptionTypeForm) error {
	var subscription *entity.SubscriptionType
	err := dao.SubscriptionType.Ctx(ctx).Where("id", form.Id).Scan(&subscription)
	if err != nil {
		return err
	}
	if subscription == nil {
		return gerror.New("Subscription not found")
	}
	subscription.Amount = form.Amount
	subscription.Days = form.Days

	_, err = dao.SubscriptionType.Ctx(ctx).Data(subscription).Where("id", form.Id).Update()

	// update cache
	err = s.InitSubscriptionSTypesToCache(ctx)
	return err
}
