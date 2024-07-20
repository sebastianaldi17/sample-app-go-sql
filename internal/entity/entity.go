package entity

import "time"

type Todo struct {
	ID         int64     `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	LastUpdate time.Time `json:"last_update" db:"last_update"`
}

type UpdateTodoRequest struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type InsertTodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
