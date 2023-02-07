package order

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"strconv"

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

func (o *sOrder) CreateOrder(ctx context.Context, input model.CreateOrderForm) (response *model.ResponseOrderForm, err error) {
	response = &model.ResponseOrderForm{}
	// check if the time is taken
	validated, err := o.ValidateTime(ctx, input)
	if err != nil {
		return
	}
	if !validated {
		err = gerror.New("time is taken or invalid")
		return
	}
	// TODO: check amount
	facilityPlace, err := service.Place().GetPlaceById(ctx, input.PlaceId)
	if err != nil {
		err = gerror.New("place not found")
		return
	}
	orderEntity := &entity.Order{}
	orderEntity.UserId = input.UserId
	orderEntity.PlaceId = input.PlaceId
	orderEntity.StartTime = gtime.NewFromStr(input.StartTime)
	orderEntity.EndTime = gtime.NewFromStr(input.EndTime)
	orderEntity.Amount = facilityPlace.Cost
	// TODO: order code
	orderEntity.OrderCode = o.GenerateOrderCode()
	orderEntity.Time = gtime.Now()

	_, err = dao.Order.Ctx(ctx).Save(orderEntity)
	response.OrderCode = orderEntity.OrderCode
	response.Amount = orderEntity.Amount
	response.StartTime = orderEntity.StartTime.String()
	response.EndTime = orderEntity.EndTime.String()
	response.PlaceId = orderEntity.PlaceId

	if err != nil {
		return
	}
	return
}

func (o *sOrder) GenerateOrderCode() string {
	// YearMonthDay + 8 digits
	return gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
}

func (o *sOrder) ValidateTime(ctx context.Context, input model.CreateOrderForm) (res bool, err error) {
	// start time should be before end time
	if gtime.NewFromStr(input.StartTime).Timestamp() >= gtime.NewFromStr(input.EndTime).Timestamp() {
		return false, nil
	}

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
	// check if user exists
	userEntity, err := service.User().GetUserByID(ctx, uint(userId))
	if err != nil {
		return
	}
	if userEntity == nil {
		err = gerror.New("user not found")
		return
	}
	err = dao.Order.Ctx(ctx).Where("user_id", userId).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GetOrdersByPlaceId(ctx context.Context, placeId int) (res []*entity.Order, err error) {
	// first 50 orders
	err = dao.Order.Ctx(ctx).Where("place_id", placeId).Scan(&res)
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

func (o *sOrder) GetAllOrders(ctx context.Context) (res []*entity.Order, err error) {
	err = dao.Order.Ctx(ctx).Scan(&res)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GetOrderByTimeRange(ctx context.Context, startTime *gtime.Time, endTime *gtime.Time) (res []*entity.Order, err error) {
	// get all orders >= start time and <= end time
	err = dao.Order.Ctx(ctx).Where("start_time >= ?", startTime).Where("end_time <= ?", endTime).Scan(&res)
	if err != nil {
		return
	}
	return
}
