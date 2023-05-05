package v1

import (
	"Gym-backend/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetWalletReq struct {
	g.Meta `path:"/wallet/wallet" method:"post" mime:"application/json" tags:"Wallet" summary:"Get wallet"`
}

type GetWalletRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Info   *model.WalletInfo `json:"info"`
}

type TopUpReq struct {
	g.Meta `path:"/wallet/top-up" method:"post" mime:"application/json" tags:"Wallet" summary:"Top up"`
	Amount float64 `json:"amount" v:"required#Please input amount"`
	CardId int     `json:"cardId" v:"required#Please input card id"`
}

type TopUpRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type WithDrawReq struct {
	g.Meta `path:"/wallet/withdraw" method:"post" mime:"application/json" tags:"Wallet" summary:"Withdraw"`
	Amount float64 `json:"amount" v:"required#Please input amount"`
	CardId int     `json:"cardId" v:"required#Please input card id"`
}

type WithDrawRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
