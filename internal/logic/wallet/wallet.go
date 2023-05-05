package wallet

import (
	"Gym-backend/internal/consts"
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sWallet struct {
}

func New() *sWallet {
	return &sWallet{}
}

func init() {
	service.RegisterWallet(New())
}

func (s *sWallet) GetWallet(ctx context.Context) (wallet *entity.Wallet, err error) {
	userId := service.Session().GetUser(ctx).Id
	err = dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
	if err != nil {
		return
	}
	return
}

func (s *sWallet) Pay(ctx context.Context, input *model.WalletPayForm) error {
	return dao.Wallet.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.Session().GetUser(ctx).Id
		var wallet *entity.Wallet
		err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
		if err != nil {
			return err
		}
		if wallet.Status != consts.WalletStatusNormal {
			return gerror.New("wallet is frozen")
		}

		if input.OrderId == 0 {
			return gerror.New("order id is 0")
		}
		var order *entity.Order
		err = dao.Order.Ctx(ctx).Where("id", input.OrderId).Scan(&order)
		if err != nil {
			return err
		}
		if order == nil {
			return gerror.New("order not exists")
		}
		if order.Status != consts.OrderStatusWaitingPayment {
			return gerror.New("order status is not waiting payment")
		}
		if order.UserId != userId {
			return gerror.New("order not belongs to user")
		}
		if order.Amount != input.Amount {
			return gerror.New("order amount not match")
		}
		if order.Amount > wallet.Amount {
			return gerror.New("wallet amount not enough")
		}
		wallet.Amount -= order.Amount
		_, err = dao.Wallet.Ctx(ctx).Where("id", wallet.Id).Update(wallet)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sWallet) PayForSubscription(ctx context.Context, input *model.WalletPayForm) error {
	return dao.Wallet.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.Session().GetUser(ctx).Id
		var wallet *entity.Wallet
		err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
		if err != nil {
			return err
		}
		if wallet.Status != consts.WalletStatusNormal {
			return gerror.New("wallet is frozen")
		}

		wallet.Amount -= input.Amount
		_, err = dao.Wallet.Ctx(ctx).Where("id", wallet.Id).Update(wallet)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sWallet) GetFullWalletInfo(ctx context.Context) (walletInfo *model.WalletInfo, err error) {
	user := service.Session().GetUser(ctx)
	userId := user.Id
	var walletEntity *entity.Wallet
	err = dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&walletEntity)
	if err != nil {
		return
	}
	var cards []*entity.WalletCard
	err = dao.WalletCard.Ctx(ctx).Where("wallet_id", walletEntity.Id).Scan(&cards)
	if err != nil {
		return
	}
	walletInfo = &model.WalletInfo{
		Wallet: walletEntity,
		Cards:  cards,
	}
	return

}

func (s *sWallet) CreateWallet(ctx context.Context) error {
	userId := service.Session().GetUser(ctx).Id
	wallet := entity.Wallet{
		UserId: userId,
		Status: 0,
	}
	_, err := dao.Wallet.Ctx(ctx).Insert(&wallet)
	return err
}

func (s *sWallet) CreateWalletForUser(ctx context.Context, userId int) error {
	if userId == 0 {
		return gerror.New("user id is 0")
	}
	wallet := entity.Wallet{
		UserId: userId,
		Status: 0,
	}
	_, err := dao.Wallet.Ctx(ctx).Insert(&wallet)
	return err
}

func (s *sWallet) SetStatus(ctx context.Context, status int) error {
	userId := service.Session().GetUser(ctx).Id
	var wallet entity.Wallet
	err := dao.Wallet.Ctx(ctx).Where("userId", userId).Scan(&wallet)
	if err != nil {
		return err
	}
	wallet.Status = status
	_, err = dao.Wallet.Ctx(ctx).Where("userId", userId).Update(&wallet)
	return err
}

func (s *sWallet) GetCardsInWallet(ctx context.Context) (cards []*entity.WalletCard, err error) {
	userId := service.Session().GetUser(ctx).Id
	var wallet *entity.Wallet
	err = dao.Wallet.Ctx(ctx).Where("userId", userId).Scan(&wallet)
	if err != nil {
		return
	}
	walletId := wallet.Id
	err = dao.WalletCard.Ctx(ctx).Where("wallet_id", walletId).Scan(&cards)
	if err != nil {
		return
	}
	return
}

func (s *sWallet) Refund(ctx context.Context, order *entity.Order) error {
	// TODO: refund fee
	return dao.Wallet.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var wallet *entity.Wallet
		var payment *entity.Payment
		user := service.Session().GetUser(ctx)
		if user.Id != order.UserId {
			return gerror.New("order not belongs to user")
		}
		err := dao.Wallet.Ctx(ctx).Where("user_id", order.UserId).Scan(&wallet)
		if err != nil {
			return err
		}
		if wallet.Status != consts.WalletStatusNormal {
			return gerror.New("wallet is frozen")
		}
		wallet.Amount += order.Amount
		_, err = dao.Wallet.Ctx(ctx).Where("id", wallet.Id).Update(wallet)
		if err != nil {
			return err
		}
		err = service.Payment().UpdatePaymentStatusByOrderId(ctx, order.Id, consts.PaymentRefund)
		if err != nil {
			return err
		}

		payment, err = service.Payment().GetPaymentByOrderId(ctx, order.Id)
		if err != nil {
			fmt.Println("get payment by order id error")
			return err
		}

		form := &model.RefundPaymentForm{
			Order:           order,
			OriginalPayment: payment,
		}
		err = service.Payment().CreatePaymentForRefund(ctx, form)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sWallet) TopUp(ctx context.Context, input *model.WalletTopUpForm) error {
	return dao.Wallet.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.Session().GetUser(ctx).Id
		var wallet *entity.Wallet
		err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
		if err != nil {
			return err
		}
		if wallet.Status != consts.WalletStatusNormal {
			return gerror.New("wallet is frozen")
		}
		var card *entity.WalletCard
		err = dao.WalletCard.Ctx(ctx).Where("id", input.CardId).Scan(&card)
		if card.WalletId != wallet.Id {
			return gerror.New("card not belongs to wallet")
		}
		if err != nil {
			return err
		}
		if card.Amount < input.Amount {
			return gerror.New("card amount not enough")
		}
		wallet.Amount += input.Amount
		_, err = dao.Wallet.Ctx(ctx).Where("id", wallet.Id).Update(wallet)
		if err != nil {
			return err
		}
		card.Amount -= input.Amount
		_, err = dao.WalletCard.Ctx(ctx).Where("id", card.Id).Update(card)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sWallet) WithDraw(ctx context.Context, input *model.WalletWithdrawForm) error {
	return dao.Wallet.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.Session().GetUser(ctx).Id
		var wallet *entity.Wallet
		err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
		if err != nil {
			return err
		}
		if wallet.Status != consts.WalletStatusNormal {
			return gerror.New("wallet is frozen")
		}
		var card *entity.WalletCard
		err = dao.WalletCard.Ctx(ctx).Where("id", input.CardId).Scan(&card)
		if card.WalletId != wallet.Id {
			return gerror.New("card not belongs to wallet")
		}
		if err != nil {
			return err
		}
		if wallet.Amount < input.Amount {
			return gerror.New("wallet amount not enough")
		}
		wallet.Amount -= input.Amount
		_, err = dao.Wallet.Ctx(ctx).Where("id", wallet.Id).Update(wallet)
		if err != nil {
			return err
		}
		card.Amount += input.Amount
		_, err = dao.WalletCard.Ctx(ctx).Where("id", card.Id).Update(card)
		if err != nil {
			return err
		}
		return nil
	})
}
