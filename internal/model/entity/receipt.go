// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Receipt is the golang structure for table receipt.
type Receipt struct {
	Id          int    `json:"id"          ` // ID
	OrderId     string `json:"orderId"     ` // Order Id
	ReceiptPath string `json:"receiptPath" ` // PDF Receipt path
}
