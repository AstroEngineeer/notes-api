package models

type Note struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user-id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
