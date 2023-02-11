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
