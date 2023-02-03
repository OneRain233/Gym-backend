package order

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"math/rand"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

type sOrder struct {
}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

func (o *sOrder) CreateOrder(ctx context.Context, input model.CreateOrderForm) error {
	// check if the time is taken
	res, err := o.ValidateTime(ctx, input)
	if err != nil || !res {
		return gerror.New("The time is taken")
	}
	var orderEntity *entity.Order
	orderEntity.UserId = input.UserId
	orderEntity.PlaceId = input.PlaceId
	orderEntity.StartTime = gtime.NewFromStr(input.StartTime)
	orderEntity.EndTime = gtime.NewFromStr(input.EndTime)
	orderEntity.Amount = input.Amount
	// TODO: order code
	orderEntity.OrderCode = string(rune(rand.Int()))
	orderEntity.PaymentType = input.PaymentType

	_, err = dao.Order.Ctx(ctx).Save(orderEntity)
	if err != nil {
		return err
	}
	return nil
}

func (o *sOrder) ValidateTime(ctx context.Context, input model.CreateOrderForm) (res bool, err error) {
	// find all orders in the same place
	var orders []*entity.Order
	// TODO: optimize the query
	err = dao.Order.Ctx(ctx).Where("place_id", input.PlaceId).Scan(&orders)

	if err != nil {
		return false, err
	}
	// check if the time is taken
	startTime := gtime.NewFromStr(input.StartTime).Timestamp()
	endTime := gtime.NewFromStr(input.EndTime).Timestamp()
	for _, order := range orders {
		tmpStartTime := order.StartTime.Timestamp()
		tmpEndTime := order.EndTime.Timestamp()
		// if the start time is in the range of an order
		if startTime >= tmpStartTime && startTime <= tmpEndTime {
			return false, nil
		}
		// if the end time is in the range of an order
		if endTime >= tmpStartTime && endTime <= tmpEndTime {
			return false, nil
		}
	}
	return true, nil

}

func (o *sOrder) GetOrdersByUserId(ctx context.Context, userId int) (res []*entity.Order, err error) {
	err = dao.Order.Ctx(ctx).Where("user_id", userId).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GetOrdersByPlaceId(ctx context.Context, placeId int) (res []*entity.Order, err error) {
	// first 50 orders
	err = dao.Order.Ctx(ctx).Where("place_id", placeId).Limit(50).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GetOrderByOrderCode(ctx context.Context, orderCode string) (res *entity.Order, err error) {
	err = dao.Order.Ctx(ctx).Where("order_code", orderCode).Scan(&res)
	if err != nil {
		return
	}
	return
}