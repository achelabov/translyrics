package models

type Article struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}
