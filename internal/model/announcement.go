package model

type AddAnnouncement struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Images  string `json:"images"`
}
