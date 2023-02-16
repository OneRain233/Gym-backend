package model

import "Gym-backend/internal/model/entity"

type BindCardForm struct {
	BankId      int     `json:"bankId"`
	CardAccount string  `json:"cardAccount"`
	Phone       string  `json:"phone"`
	Amount      float64 `json:"amount"` // simulated amount
}

type CardPayForm struct {
	CardId  int     `json:"cardId"`
	Amount  float64 `json:"amount"`
	OrderId int     `json:"orderId"`
}

type CardRechargeForm struct {
	CardId  int     `json:"cardId"`
	Amount  float64 `json:"amount"`
	OrderId int     `json:"orderId"`
}

type WalletInfo struct {
	Wallet *entity.Wallet       `json:"wallet"`
	Cards  []*entity.WalletCard `json:"cards"`
}

type WalletPayForm struct {
	Amount  float64 `json:"amount"`
	OrderId int     `json:"orderId"`
}

type WalletTopUpForm struct {
	Amount float64 `json:"amount"`
	CardId int     `json:"cardId"`
}

type WalletWithdrawForm struct {
	Amount float64 `json:"amount"`
	CardId int     `json:"cardId"`
}
