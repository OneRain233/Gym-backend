package model

type BindCardForm struct {
	BankId      int    `json:"bankId"`
	CardAccount string `json:"cardAccount"`
	Phone       string `json:"phone"`
	Amount      string `json:"amount"` // simulated amount
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
