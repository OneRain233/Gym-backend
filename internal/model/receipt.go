package model

type ReceiptInfo struct {
	Facility    string `json:"facility"`
	Activity    string `json:"activity"`
	StartTime   string `json:"startTime"`
	LastingTime string `json:"lastingTime"`
	Amount      string `json:"amount"`
}
