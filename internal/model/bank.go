package model

type AddBankForm struct {
	Name string `json:"name"`
}

type UpdateBankForm struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
