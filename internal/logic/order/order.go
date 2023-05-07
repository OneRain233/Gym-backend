package order

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"Gym-backend/utility/receipt"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"

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

func (o *sOrder) PayByCash(ctx context.Context, orderCode string) error {
	var order *entity.Order
	err := dao.Order.Ctx(ctx).Where("order_code", orderCode).Scan(&order)
	if err != nil {
		return err
	}
	paymentForm := &model.CreatePaymentForm{
		OrderCode:   orderCode,
		PaymentType: consts.PaymentTypeCash,
	}
	_, err = service.Payment().CreatePayment(ctx, paymentForm)
	if err != nil {
		return err
	}
	return nil
}

func (o *sOrder) CalculateAmount(ctx context.Context, input model.CreateOrderForm) (amount float64, discount float64, err error) {
	unit := 0
	startTime := input.StartTime
	endTime := input.EndTime
	startTimeTime, err := gtime.StrToTime(startTime)
	if err != nil {
		return
	}
	endTimeTime, err := gtime.StrToTime(endTime)
	if err != nil {
		return
	}
	duration := endTimeTime.Sub(startTimeTime)
	userId := service.Session().GetUser(ctx).Id
	subscription, err := service.Subscription().GetSubscriptionByUserId(ctx, userId)

	var subscriptionType *entity.SubscriptionType
	if subscription != nil {
		subscriptionType, err = service.Subscription().GetSubscriptionTypeById(ctx, subscription.SubscriptionType)
		if err != nil {
			return
		}
	} else {
		subscriptionType = &entity.SubscriptionType{}
	}

	//place := &entity.FacilityPlace{}
	place, err := service.Place().GetPlaceById(ctx, input.PlaceId)
	if err != nil {
		return
	}

	if duration.Hours() < 0 {
		err = gerror.New("invalid time")
		return
	}
	if duration.Hours() < 1 {
		unit = 1
		return
	}
	unit = unit + int(duration.Hours()/1)

	if subscriptionType.Name == "monthly" {
		amount = place.Cost * float64(unit) * 0.8
		discount = 0.8
		return
	} else if subscriptionType.Name == "yearly" {
		amount = place.Cost * float64(unit) * 0.7
		discount = 0.7
		return
	} else {
		amount = place.Cost * float64(unit)
		discount = 1
		return
	}
}

func (o *sOrder) CreateRegularWeeklyOrder(ctx context.Context, input model.CreateRegularOrderFormWeekly) (response *model.RegularOrderResponseForm, err error) {
	response = &model.RegularOrderResponseForm{}
	response.Orders = make([]*model.ResponseOrderForm, 0)
	startDay := gtime.NewFromStr(input.StartDay)
	weekCount := input.WeekCount
	parentOrderCode := o.GenerateParentOrderCode()
	// calculate every session date
	var sessionDate []*gtime.Time
	for i := 0; i < weekCount; i++ {
		sessionDate = append(sessionDate, startDay.AddDate(0, 0, 7*i))
	}
	// check if the time is taken
	startTime := gtime.NewFromStr(input.SessionStartTime)
	endTime := gtime.NewFromStr(input.SessionEndTime)
	for _, date := range sessionDate {
		newInput := model.CreateOrderForm{
			UserId:    input.UserId,
			PlaceId:   input.PlaceId,
			StartTime: date.Format("Y-m-d") + " " + startTime.Format("H:i:s"),
			EndTime:   date.Format("Y-m-d") + " " + endTime.Format("H:i:s"),
		}
		validated, err := o.ValidateTime(ctx, newInput)
		if err != nil {
			return response, err
		}
		if !validated {
			err = gerror.New("time is taken or invalid")
			return response, err
		}
	}

	// calculate amount
	amountCnt, discount, err := o.CalculateAmount(ctx, model.CreateOrderForm{
		UserId:    input.UserId,
		PlaceId:   input.PlaceId,
		StartTime: sessionDate[0].Format("Y-m-d") + " " + startTime.Format("H:i:s"),
		EndTime:   sessionDate[0].Format("Y-m-d") + " " + endTime.Format("H:i:s"),
	})
	if err != nil {
		return
	}
	amount := amountCnt * float64(weekCount)
	// create order for every session
	for _, date := range sessionDate {
		newInput := model.CreateOrderForm{
			UserId:    input.UserId,
			PlaceId:   input.PlaceId,
			StartTime: date.Format("Y-m-d") + " " + startTime.Format("H:i:s"),
			EndTime:   date.Format("Y-m-d") + " " + endTime.Format("H:i:s"),
		}
		orderEntity := &entity.Order{
			Time:            gtime.Now(),
			OrderCode:       o.GenerateOrderCode(),
			UserId:          newInput.UserId,
			PlaceId:         newInput.PlaceId,
			StartTime:       gtime.NewFromStr(newInput.StartTime),
			EndTime:         gtime.NewFromStr(newInput.EndTime),
			Amount:          amountCnt,
			Status:          consts.OrderStatusWaitingPayment,
			Discount:        discount,
			ParentOrderCode: parentOrderCode,
		}
		_, err = dao.Order.Ctx(ctx).Insert(orderEntity)
		if err != nil {
			return response, err
		}
		response.Orders = append(response.Orders, &model.ResponseOrderForm{
			Amount:    amountCnt,
			OrderCode: orderEntity.OrderCode,
			StartTime: newInput.StartTime,
			EndTime:   newInput.EndTime,
			PlaceId:   newInput.PlaceId,
		})
	}
	response.Amount = amount
	response.OrderCode = parentOrderCode
	return

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

	amountCnt, discount, err := o.CalculateAmount(ctx, input)
	if err != nil {
		return
	}
	orderEntity := &entity.Order{
		UserId:    input.UserId,
		PlaceId:   input.PlaceId,
		StartTime: gtime.NewFromStr(input.StartTime),
		EndTime:   gtime.NewFromStr(input.EndTime),
		Amount:    amountCnt,
		Status:    consts.OrderStatusWaitingPayment,
		OrderCode: o.GenerateOrderCode(),
		Time:      gtime.Now(),
		Discount:  discount,
	}
	//orderEntity.UserId = input.UserId
	//orderEntity.PlaceId = input.PlaceId
	//orderEntity.StartTime = gtime.NewFromStr(input.StartTime)
	//orderEntity.EndTime = gtime.NewFromStr(input.EndTime)
	//orderEntity.Amount = amountCnt
	//orderEntity.Status = consts.OrderStatusWaitingPayment
	//orderEntity.OrderCode = o.GenerateOrderCode()
	//orderEntity.Time = gtime.Now()
	//orderEntity.Discount = discount

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
	orderCode := gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
	if len(orderCode) != 17 {
		return o.GenerateOrderCode()
	}
	return orderCode
}

func (o *sOrder) GenerateParentOrderCode() string {
	// "2" + YearMonthDay + 8 digits
	orderCode := "2" + gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
	if len(orderCode) != 18 {
		return o.GenerateParentOrderCode()
	}
	return orderCode
	//return "2" + gtime.Now().Format("Ymd") + strconv.Itoa(gtime.Now().Nanosecond())
}

func (o *sOrder) ValidateTime(ctx context.Context, input model.CreateOrderForm) (res bool, err error) {
	// check if we have open
	OpenTimeEntity, err := service.Config().GetConfigByKey(ctx, consts.OpenTime)
	if err != nil {
		return false, err
	}
	CloseTimeEntity, err := service.Config().GetConfigByKey(ctx, consts.CloseTime)
	if err != nil {
		return false, err
	}

	StartTime := gtime.NewFromStr(input.StartTime)
	EndTime := gtime.NewFromStr(input.EndTime)
	OpenTime, err := gtime.StrToTime(StartTime.Format("Y-m-d") + " " + OpenTimeEntity.Value)
	CloseTime, err := gtime.StrToTime(EndTime.Format("Y-m-d") + " " + CloseTimeEntity.Value)

	// check if the time is in open time
	if StartTime.Timestamp() < OpenTime.Timestamp() || EndTime.Timestamp() > CloseTime.Timestamp() {
		return false, gerror.New("invalid time(we are not open now)")
	}

	if gtime.NewFromStr(input.StartTime).Timestamp() < OpenTime.Timestamp() || gtime.NewFromStr(input.EndTime).Timestamp() > CloseTime.Timestamp() {
		return false, gerror.New("invalid time(we are not open now)")
	}

	// check if the time is end with 00 15 30 45
	if gtime.NewFromStr(input.StartTime).Minute() != 0 && gtime.NewFromStr(input.StartTime).Minute() != 15 && gtime.NewFromStr(input.StartTime).Minute() != 30 && gtime.NewFromStr(input.StartTime).Minute() != 45 {
		return false, gerror.New("invalid time(start time does not end with 00 15 30 45)")
	}
	if gtime.NewFromStr(input.EndTime).Minute() != 0 && gtime.NewFromStr(input.EndTime).Minute() != 15 && gtime.NewFromStr(input.EndTime).Minute() != 30 && gtime.NewFromStr(input.EndTime).Minute() != 45 {
		return false, gerror.New("invalid time(end time does not end with 00 15 30 45)")
	}

	// start time should be before end time
	if gtime.NewFromStr(input.StartTime).Timestamp() >= gtime.NewFromStr(input.EndTime).Timestamp() {
		return false, gerror.New("invalid time(start time should be before end time)")
	}

	// check if the time before now
	if gtime.NewFromStr(input.StartTime).Timestamp() < gtime.Now().Timestamp() {
		return false, gerror.New("invalid time(start time should be after now)")
	}

	// check if startTime and endTime in the same day
	if gtime.NewFromStr(input.StartTime).Day() != gtime.NewFromStr(input.EndTime).Day() {
		return false, gerror.New("invalid time(start time and end time should be in the same day)")
	}

	// find all orders in the same place
	var orders []*entity.Order
	// TODO: optimize the query for time
	queryStartTime := gtime.NewFromStr(input.StartTime).Add(time.Duration(-24) * time.Hour) // 24 hour before
	queryEndTime := gtime.NewFromStr(input.EndTime).Add(time.Duration(24) * time.Hour)      // 24 hour after
	err = dao.Order.Ctx(ctx).Where("place_id = ? AND start_time >= ? AND end_time <= ? AND status = ?", input.PlaceId, queryStartTime, queryEndTime, consts.OrderStatusPending).Scan(&orders)
	//err = dao.Order.Ctx(ctx).Where("place_id", input.PlaceId).Where("status", consts.OrderStatusPending).Scan(&orders)
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
			return false, gerror.New("invalid time(start time is in the range of an order)")
		}
		// if the end time is in the range of an order
		if endTime >= tmpStartTime && endTime <= tmpEndTime {
			return false, gerror.New("invalid time(end time is in the range of an order)")
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

func (o *sOrder) GetRegularOrdersByOrderCode(ctx context.Context, orderCode string) (res []*entity.Order, err error) {
	res = []*entity.Order{}
	var orders []*entity.Order
	err = dao.Order.Ctx(ctx).Where("parent_order_code", orderCode).Scan(&orders)
	if err != nil {
		return
	}
	if len(orders) == 0 {
		return nil, gerror.New("order not found")
	}
	res = append(res, orders...)
	return
}

func (o *sOrder) GetOrderByOrderCode(ctx context.Context, orderCode string) (res *entity.Order, err error) {
	res = &entity.Order{}
	err = dao.Order.Ctx(ctx).Where("order_code", orderCode).Scan(&res)
	if res == nil {
		return nil, gerror.New("order not found")
	}
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GetRegularOrderByParentOrderCode(ctx context.Context, parentOrderCode string) (res *model.RegularOrderInfoResponseForm, err error) {
	res = &model.RegularOrderInfoResponseForm{}
	var orders []*entity.Order
	err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().ParentOrderCode, parentOrderCode).Scan(&orders)
	if err != nil {
		return
	}

	res.OrderCode = parentOrderCode
	res.Amount = 0
	for _, order := range orders {
		res.Amount += order.Amount
	}
	res.Orders = orders
	return
}

func (o *sOrder) GetAllOrders(ctx context.Context, pagination *model.Pagination) (res []*entity.Order, err error) {
	if pagination.Page == 0 || pagination.Limit == 0 {
		err = dao.Order.Ctx(ctx).Scan(&res)
		return
	} else {
		offset := (pagination.Page - 1) * pagination.Limit
		err = dao.Order.Ctx(ctx).Limit(pagination.Limit).Offset(offset).Scan(&res)
		return
	}
}

func (o *sOrder) GetRefundedOrder(ctx context.Context) (res []*entity.Order, err error) {
	err = dao.Order.Ctx(ctx).Where("status", consts.OrderStatusRefunded).Scan(&res)
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

func (o *sOrder) GenerateOrderReceipt(ctx context.Context, orderCode string) (path string, err error) {
	// Test Cases:
	// 1. If there is no such order 		pass
	// 2. If the order is not paid			pass
	// 3. If there is already a receipt		pass
	// 4. If the order is paid and there is no receipt	pass
	// 5. If the order is paid and there is a receipt	pass

	// TODO: regenerate?
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return "", err
	}
	if order == nil {
		return "", gerror.New("order not found")
	}
	receiptEntity, err := service.Receipt().GetReceiptByOrderId(ctx, order.Id)
	if err != nil {
		return "", err
	}
	if receiptEntity != nil {
		// check file exists?
		if gfile.Exists(receiptEntity.ReceiptPath) {
			return receiptEntity.ReceiptPath, nil
		}
		// delete the receipt in db
		err = service.Receipt().DeleteReceipt(ctx, receiptEntity.Id)
		if err != nil {
			return "", err
		}
	}
	// check if the order is paid
	payment, err := service.Payment().GetPaymentByOrderCode(ctx, orderCode)
	if err != nil {
		return "", err
	}
	if payment == nil {
		return "", gerror.New("payment not found")
	}
	if payment.Status != consts.PaymentSuccess {
		return "", gerror.New("payment not successful")
	}

	if order.Status != consts.OrderStatusDone {
		return "", gerror.New("order not done")
	}

	// generate order code qr code
	userId := service.Session().GetUser(ctx).Id
	qeFilename := strconv.Itoa(int(userId)) + orderCode + ".png"

	// json qrContent
	// TODO: qr code just contain the order code
	qrContent := map[string]interface{}{
		"orderCode":     order.OrderCode,
		"user":          userId,
		"generatedTime": gtime.Now().String(),
	}
	sign := o.GenerateQrSignature(qrContent)
	qrContent["sign"] = sign
	qrFilePath, err := receipt.GenerateQRCode(qeFilename, gjson.MustEncodeString(qrContent))
	if err != nil {
		return "", err
	}
	// generate order pdf
	pdfFilename := strconv.Itoa(int(userId)) + orderCode + ".pdf"

	place, err := service.Place().GetPlaceById(ctx, order.PlaceId)
	if err != nil {
		return "", err
	}
	facility, err := service.Facility().GetFacilityById(ctx, place.FacilityId)
	if err != nil {
		return "", err
	}

	pdfContent := &model.ReceiptInfo{
		Facility:    facility.Facility.Name + " " + place.Name,
		Activity:    "Rent",
		StartTime:   order.StartTime.String(),
		LastingTime: order.EndTime.String(),
		Amount:      strconv.Itoa(int(order.Amount)),
	}

	path, err = receipt.GenerateReceiptPDF(pdfFilename, pdfContent, qrFilePath)
	if err != nil {
		return "", err
	}
	// recode in db
	form := model.CreateReceiptForm{
		OrderCode:   orderCode,
		ReceiptPath: path,
		OrderId:     order.Id,
	}
	err = service.Receipt().AddReceipt(ctx, form)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) RefundOrder(ctx context.Context, orderCode string) (err error) {
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}
	if order.Status != consts.OrderStatusPending {
		err = gerror.New("order not paid")
		return
	}
	user := service.Session().GetUser(ctx)
	if user.Id != order.UserId {
		err = gerror.New("user not authorized")
		return
	}
	payment, err := service.Payment().GetPaymentByOrderId(ctx, order.Id)
	if err != nil {
		return
	}
	if payment == nil {
		err = gerror.New("payment not found, please contact the manager")
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = service.Wallet().Refund(ctx, order)
		if err != nil {
			return err
		}
		order.Status = consts.OrderStatusRefunded
		_, err = dao.Order.Ctx(ctx).Where("id", order.Id).Data(order).Update()
		if err != nil {
			return err
		}
		return nil
	})

}

func (o *sOrder) UpdateOrderStatus(ctx context.Context, orderCode string, status int) (err error) {
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}
	order.Status = status
	_, err = dao.Order.Ctx(ctx).Where("id", order.Id).Data(order).Update()
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GenerateQrSignature(qrContent map[string]interface{}) string {
	// TODO: Add a salt
	res := gmd5.MustEncryptString(gjson.MustEncodeString(qrContent))
	return res
}

func (o *sOrder) CheckSignature(qrContent map[string]interface{}, sign string) bool {
	originalQrContent := map[string]interface{}{
		"orderCode":     qrContent["orderCode"],
		"user":          qrContent["user"],
		"generatedTime": qrContent["generatedTime"],
	}
	originalSign := o.GenerateQrSignature(originalQrContent)
	if originalSign != sign {
		return false
	}
	return true
}

func (o *sOrder) StartOrder(ctx context.Context, orderCode string) (err error) {
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return
	}
	if order == nil {
		return gerror.New("order not found")
	}
	switch order.Status {
	case consts.OrderStatusWaitingPayment:
		return gerror.New("order not paid")

	case consts.OrderStatusDoing:
		return gerror.New("order already started")

	case consts.OrderStatusCancelled:
		return gerror.New("order cancelled")

	case consts.OrderStatusDone:
		return gerror.New("order already done")
	}
	order.Status = consts.OrderStatusDoing
	_, err = dao.Order.Ctx(ctx).Where("order_code", orderCode).Update(order)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) EndOrder(ctx context.Context, orderCode string) (err error) {
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return
	}
	if order == nil {
		return gerror.New("order not found")
	}
	switch order.Status {
	case consts.OrderStatusWaitingPayment:
		return gerror.New("order not paid")

	case consts.OrderStatusCancelled:
		return gerror.New("order cancelled")

	case consts.OrderStatusDone:
		return gerror.New("order already done")

	case consts.OrderStatusPending:
		return gerror.New("order not started")

	}
	order.Status = consts.OrderStatusDone
	_, err = dao.Order.Ctx(ctx).Where("order_code", orderCode).Update(order)
	if err != nil {
		return
	}

	return
}

func (o *sOrder) CancelOrder(ctx context.Context, orderCode string) (err error) {
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return
	}
	if order == nil {
		return gerror.New("order not found")
	}
	switch order.Status {
	case consts.OrderStatusCancelled:
		return gerror.New("order already cancelled")

	case consts.OrderStatusDone:
		return gerror.New("order already done")

	case consts.OrderStatusDoing:
		return gerror.New("order already started")
	}
	order.Status = consts.OrderStatusCancelled
	_, err = dao.Order.Ctx(ctx).Where("order_code", orderCode).Update(order)
	if err != nil {
		return
	}
	return
}

func (o *sOrder) GetDailyOrderIncome(ctx context.Context, date *gtime.Time) (income float64, err error) {
	startDate := date.Format("Y-m-d 00:00:00")
	endDate := date.Format("Y-m-d 23:59:59")
	var orders []*entity.Order
	err = dao.Order.Ctx(ctx).Where("status", consts.PaymentSuccess).Where("time BETWEEN ? AND ?", startDate, endDate).Scan(&orders)
	if err != nil {
		return
	}
	for _, order := range orders {
		income += order.Amount
	}
	return
}

func (o *sOrder) GetWeeklyOrderIncome(ctx context.Context, data *gtime.Time) (income float64, err error) {
	startDate := data.Format("Y-m-d 00:00:00")
	endDate := data.AddDate(0, 0, 7).Format("Y-m-d 23:59:59")
	var orders []*entity.Order
	err = dao.Order.Ctx(ctx).Where("status", consts.PaymentSuccess).Where("time BETWEEN ? AND ?", startDate, endDate).Scan(&orders)
	if err != nil {
		return
	}
	for _, order := range orders {
		income += order.Amount
	}
	return
}

func (o *sOrder) CheckExpiredOrder(ctx context.Context) (err error) {
	// using order -> done
	var usingOrder []*entity.Order
	err = dao.Order.Ctx(ctx).Where("status", consts.OrderStatusDoing).Scan(&usingOrder)
	if err != nil {
		return
	}
	fmt.Println("using order", usingOrder)
	// bug: didn't send email
	for _, order := range usingOrder {
		if order.EndTime.Before(gtime.Now()) {
			fmt.Println("end order")
			err = o.EndOrder(ctx, order.OrderCode)
			if err != nil {
				g.Log().Error(ctx, err.Error())
				continue
			}
			fmt.Println("send receipt")
			err = service.Receipt().SendReceiptToUser(ctx, order.OrderCode)
			if err != nil {
				g.Log().Error(ctx, err.Error())
				continue
			}
		}
	}
	// pending order -> start order
	// TODO: I think it should refund the money
	var pendingOrder []*entity.Order
	err = dao.Order.Ctx(ctx).Where("status", consts.OrderStatusPending).Scan(&pendingOrder)
	if err != nil {
		return
	}
	for _, order := range pendingOrder {
		if order.StartTime.Timestamp() <= gtime.Now().Timestamp() {
			// refund
			err = o.UpdateOrderStatus(ctx, order.OrderCode, consts.OrderStatusCancelled)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// unpaid order
	var unpaidOrder []*entity.Order
	err = dao.Order.Ctx(ctx).Where("status", consts.OrderStatusWaitingPayment).Scan(&unpaidOrder)
	if err != nil {
		return
	}
	for _, order := range unpaidOrder {
		if order.StartTime.Before(gtime.Now()) {
			// update order status to cancelled
			err = o.CancelOrder(ctx, order.OrderCode)
			if err != nil {
				return
			}
		}
	}
	return
}

func (o *sOrder) GetOrderByStatus(ctx context.Context, status int) (order []*entity.Order, err error) {
	user := service.Session().GetUser(ctx)
	userId := user.Id
	err = dao.Order.Ctx(ctx).Where("status", status).Where("user_id", userId).Scan(&order)
	if err != nil {
		return
	}
	return
}
