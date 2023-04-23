package model

type AddCourseForm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Video       string `json:"video"`
	Image       string `json:"image"`
}
