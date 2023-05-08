package model

type CreateSubscriptionForm struct {
	UserId      int
	Type        int
	PaymentType int
	CardId      int
}

type UpdateSubscriptionTypeForm struct {
	Days   int
	Amount float64
	Id     int
}
