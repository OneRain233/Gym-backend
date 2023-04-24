package receipt

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"Gym-backend/utility/mail"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sReceipt struct {
}

func init() {
	service.RegisterReceipt(New())
}

func New() *sReceipt {
	return &sReceipt{}
}

func (r *sReceipt) GetReceiptByOrderId(ctx context.Context, orderId int) (receipt *entity.Receipt, err error) {
	receipt = &entity.Receipt{}
	// check count
	cnt, err := dao.Receipt.Ctx(ctx).Where("order_id", orderId).Count()
	if err != nil {
		return
	}
	if cnt == 0 {
		return nil, nil
	}

	err = dao.Receipt.Ctx(ctx).Where("order_id", orderId).Scan(receipt)
	if err != nil {
		return
	}
	return
}

func (r *sReceipt) AddReceipt(ctx context.Context, input model.CreateReceiptForm) (err error) {
	receiptEntity := &entity.Receipt{}
	receiptEntity.ReceiptPath = input.ReceiptPath
	receiptEntity.OrderId = input.OrderId
	_, err = dao.Receipt.Ctx(ctx).Save(receiptEntity)
	return
}

func (r *sReceipt) SendReceiptToUser(ctx context.Context, orderCode string) (err error) {
	order, err := service.Order().GetOrderByOrderCode(ctx, orderCode)
	if err != nil {
		return
	}
	if order == nil {
		err = gerror.New("order not found")
		return
	}
	receipt, err := r.GetReceiptByOrderId(ctx, order.Id)
	if err != nil {
		return
	}
	if receipt == nil {
		// generate receipt
		_, err := service.Order().GenerateOrderReceipt(ctx, order.OrderCode)
		if err != nil {
			return err
		}
	}
	user, err := service.User().GetUserById(ctx, uint(order.UserId))
	if err != nil {
		return
	}
	if user == nil {
		return gerror.New("user not found")
	}

	// send emil
	g.Log().Info(ctx, "sending email to user")
	err = mail.SendEmailWithAttachment(user.Email, "Receipt Your receipt is attached", receipt.ReceiptPath)
	return
}

func (r *sReceipt) DeleteReceipt(ctx context.Context, id int) error {
	_, err := dao.Receipt.Ctx(ctx).Where("id", id).Delete()
	return err
}
