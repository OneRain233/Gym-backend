package v1

import "github.com/gogf/gf/v2/frame/g"

type BindCardReq struct {
	g.Meta      `path:"/card/bind" tags:"Card" method:"post" summary:"Bind Card"`
	BankId      int    `json:"bank_id" v:"required#Please input bank_id"`
	CardAccount string `json:"card_account" v:"required#Please input card_account"`
	Phone       string `json:"phone" v:"required#Please input phone"`
}

type BindCardRes struct {
	g.Meta `mime:"application/json" example:"string"`
}

type DeleteCardReq struct {
	g.Meta `path:"/card/delete" tags:"Card" method:"post" summary:"Delete Card"`
	Id     int `json:"id" v:"required#Please input id"`
}

type DeleteCardRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
