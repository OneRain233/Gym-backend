package card

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/text/gregex"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sCard struct{}

func New() *sCard {
	return &sCard{}
}

func init() {
	service.RegisterCard(New())
}

func (s *sCard) GetCard(ctx context.Context) (card *entity.WalletCard, err error) {
	userId := service.Session().GetUser(ctx)
	err = dao.WalletCard.Ctx(ctx).Where("user_id", userId).Scan(&card)
	if err != nil {
		return
	}
	return
}

func (s *sCard) ValidateCard(ctx context.Context, input *model.BindCardForm) error {
	cnt, err := dao.WalletCard.Ctx(ctx).Where("card_account", input.CardAccount).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New("card account already exists")
	}

	// check format of card account
	// all card account should be digit
	if !gregex.IsMatchString(`^\d+$`, input.CardAccount) {
		return gerror.New("card account should be digit")
	}

	return nil
}

func (s *sCard) ValidatePhone(ctx context.Context, input *model.BindCardForm) error {
	if !gregex.IsMatchString(`^1[3-9]\d{9}$`, input.Phone) {
		return gerror.New("phone format is not correct")
	}
	// check length of phone
	if len(input.Phone) != 11 {
		return gerror.New("phone length is not correct")
	}
	return nil
}

func (s *sCard) BindCard(ctx context.Context, input *model.BindCardForm) error {
	userId := service.Session().GetUser(ctx).Id
	var wallet *entity.Wallet
	err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
	if err != nil {
		return err
	}
	// validate card
	err = s.ValidateCard(ctx, input)
	if err != nil {
		return err
	}
	err = s.ValidatePhone(ctx, input)
	if err != nil {
		return err
	}
	// check if bank id exists
	var bank *entity.Bank
	err = dao.Bank.Ctx(ctx).Where("id", input.BankId).Scan(&bank)
	if err != nil {
		return err
	}
	if bank == nil {
		return gerror.New("bank id not exists")
	}
	card := &entity.WalletCard{
		BankId:      input.BankId,
		WalletId:    wallet.Id,
		CardAccount: input.CardAccount,
		Phone:       input.Phone,
		Amount:      0.0,
	}
	_, err = dao.WalletCard.Ctx(ctx).Insert(card)
	return err
}

func (s *sCard) GetCardsByUserId(ctx context.Context, userId int) (cards []*entity.WalletCard, err error) {
	err = dao.WalletCard.Ctx(ctx).Where("user_id", userId).Scan(&cards)
	if err != nil {
		return
	}
	return
}

func (s *sCard) GetCardById(ctx context.Context, cardId int) (card *entity.WalletCard, err error) {
	err = dao.WalletCard.Ctx(ctx).Where("id", cardId).Scan(&card)
	if err != nil {
		return
	}
	return
}

func (s *sCard) GetCardsCountByUserId(ctx context.Context, userId int) (count int, err error) {
	count, err = dao.WalletCard.Ctx(ctx).Where("user_id", userId).Count()
	if err != nil {
		return
	}
	return
}

func (s *sCard) Pay(ctx context.Context, input *model.CardPayForm) error {
	return dao.WalletCard.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.Session().GetUser(ctx).Id
		var wallet *entity.Wallet
		err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
		var card *entity.WalletCard
		err = dao.WalletCard.Ctx(ctx).Where("id", input.CardId).Scan(&card)
		if err != nil {
			return err
		}

		// validate order id
		if input.OrderId == 0 {
			return gerror.New("order id not exists")
		}
		var order *entity.Order
		err = dao.Order.Ctx(ctx).Where("id", input.OrderId).Scan(&order)
		if err != nil {
			return err
		}
		if order == nil {
			return gerror.New("order not exists")
		}
		if order.UserId != userId {
			return gerror.New("order not belongs to you")
		}
		// validate card
		if card == nil {
			return gerror.New("card not exists")
		}
		if card.WalletId != wallet.Id {
			return gerror.New("card not belongs to you")
		}

		if order.Amount != input.Amount {
			return gerror.New("amount not match")
		}

		// validate amount
		if card.Amount < input.Amount {
			return gerror.New("amount not enough")
		}
		card.Amount -= input.Amount
		_, err = dao.WalletCard.Ctx(ctx).Data(card).Save()
		if err != nil {
			return err
		}

		return nil
	})

}

func (s *sCard) PayForSubscription(ctx context.Context, input *model.CardPayForm) error {
	return dao.WalletCard.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId := service.Session().GetUser(ctx).Id
		var wallet *entity.Wallet
		err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
		var card *entity.WalletCard
		err = dao.WalletCard.Ctx(ctx).Where("id", input.CardId).Scan(&card)
		if err != nil {
			return err
		}

		// validate card
		if card == nil {
			return gerror.New("card not exists")
		}
		if card.WalletId != wallet.Id {
			return gerror.New("card not belongs to you")
		}

		if card.Amount < input.Amount {
			return gerror.New("amount not enough")
		}

		card.Amount -= input.Amount
		_, err = dao.WalletCard.Ctx(ctx).Data(card).Save()
		if err != nil {
			return err
		}

		return nil
	})

}

func (s *sCard) Recharge(ctx context.Context, input *model.CardRechargeForm) error {
	userId := service.Session().GetUser(ctx).Id
	var wallet *entity.Wallet
	err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
	var card *entity.WalletCard
	err = dao.WalletCard.Ctx(ctx).Where("id", input.CardId).Scan(&card)
	if err != nil {
		return err
	}
	if card == nil {
		return gerror.New("card not exists")
	}
	if card.WalletId != wallet.Id {
		return gerror.New("card not belongs to you")
	}
	card.Amount += input.Amount
	_, err = dao.WalletCard.Ctx(ctx).Where("id", input.CardId).Update(&card)
	if err != nil {
		return err
	}
	return nil
}

func (s *sCard) DeleteCard(ctx context.Context, cardId int) error {
	userId := service.Session().GetUser(ctx).Id
	var wallet *entity.Wallet
	err := dao.Wallet.Ctx(ctx).Where("user_id", userId).Scan(&wallet)
	if err != nil {
		return err
	}
	var card *entity.WalletCard
	err = dao.WalletCard.Ctx(ctx).Where("id", cardId).Scan(&card)
	if err != nil {
		return err
	}
	if card == nil {
		return gerror.New("card not exists")
	}
	if card.WalletId != wallet.Id {
		return gerror.New("card not belongs to you")
	}
	_, err = dao.WalletCard.Ctx(ctx).Where("id", cardId).Delete()
	if err != nil {
		return err
	}
	return nil
}
