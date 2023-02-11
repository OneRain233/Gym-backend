package wallet

import (
	"Gym-backend/internal/dao"
	"Gym-backend/internal/model"
	"Gym-backend/internal/model/entity"
	"Gym-backend/internal/service"
	"context"

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
