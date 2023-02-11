package order

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"Gym-backend/utility/receipt"
	"context"
	"strconv"

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
	if res == nil {
		return nil, nil
	}
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

func (o *sOrder) GenerateOrderReceipt(ctx context.Context, orderCode string) (path string, err error) {
	// Test Cases:
	// 1. If there is no such order 		pass
	// 2. If the order is not paid			pass
	// 3. If there is already a receipt		pass
	// 4. If the order is paid and there is no receipt	pass
	// 5. If the order is paid and there is a receipt	pass

	// TODO: regenerate?
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
	// check if there is a receipt in db
	order, err := o.GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return "", err
	}
	receiptEntity, err := service.Receipt().GetReceiptByOrderCode(ctx, order.Id)
	if err != nil {
		return "", err
	}
	if receiptEntity != nil {
		return receiptEntity.ReceiptPath, nil
	}
	// generate order code qr code
	userId := service.Session().GetUser(ctx).Id
	qeFilename := strconv.Itoa(int(userId)) + orderCode + ".png"

	// json qrContent
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

	//pdfContent := "Order Code: " + orderCode + "\n" +
	//	"Start Time: " + order.StartTime.String() + "\n" +
	//	"End Time: " + order.EndTime.String() + "\n" +
	//	"Place: " + place.Name + "\n" +
	//	"Amount: " + strconv.Itoa(int(order.Amount)) + "\n"
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
