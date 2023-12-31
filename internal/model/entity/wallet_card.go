// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WalletCard is the golang structure for table wallet_card.
type WalletCard struct {
	Id          int     `json:"id"          ` // ID
	BankId      int     `json:"bankId"      ` // Bank ID of the card
	WalletId    int     `json:"walletId"    ` // Wallet ID of the card
	CardAccount string  `json:"cardAccount" ` // Card number
	Phone       string  `json:"phone"       ` // Phone of the card
	Amount      float64 `json:"amount"      ` // Money amount
}
