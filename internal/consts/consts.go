package consts

const (
	ContextKey = "ContextKey"

	PaymentTypeCard   = 1
	PaymentTypeWallet = 2
	PaymentTypeRefund = 3

	RoleAdmin      = 1
	RoleNormalUser = 0
	RoleManager    = 2

	UserKey = "SessionUser"

	WalletStatusNormal = 0
	WalletStatusFrozen = 1

	PaymentSuccess = 1
	PaymentFail    = 0
	PaymentPending = 2
	PaymentRefund  = 3

	OrderStatusWaitingPayment = 0 // waiting for payment
	OrderStatusPending        = 1 // waiting for use
	OrderStatusDoing          = 2 // using
	OrderStatusDone           = 3 // done
	OrderStatusCancelled      = 4 // canceled
	OrderStatusRefunded       = 5 // refunded

	PaymentNotFound = "payment not found"

	OpenTime  = "OpenTime"
	CloseTime = "CloseTime"
)

const (
	GenderMale   = 1
	GenderFemale = 2
	GenderOther  = 3
)
