package model

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
