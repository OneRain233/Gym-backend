package receipt

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
)

type sReceipt struct {
}

func init() {
	service.RegisterReceipt(New())
}

func New() *sReceipt {
	return &sReceipt{}
}

func (r *sReceipt) GetReceiptByOrderCode(ctx context.Context, orderId int) (receipt *entity.Receipt, err error) {
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
