// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Payment is the golang structure for table payment.
type Payment struct {
	Id          int         `json:"id"          ` // ID
	WalletId    int         `json:"walletId"    ` // Wallet ID
	OrderId     int         `json:"orderId"     ` // Order ID
	PaymentCode string      `json:"paymentCode" ` // Payment code
	Time        *gtime.Time `json:"time"        ` // Payment Time
	Amount      float64     `json:"amount"      ` // Payment Amount
	PaymentType int         `json:"paymentType" ` // Payment type
	Status      int         `json:"status"      ` // Payment status
}
