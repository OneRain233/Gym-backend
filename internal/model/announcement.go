package model

type AddAnnouncement struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

type ModifyAnnouncement struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}
