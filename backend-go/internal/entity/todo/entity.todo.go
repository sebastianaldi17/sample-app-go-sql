package todo

import "time"

type Todo struct {
	ID         int64     `json:"id" db:"id"`
	AuthorID   int64     `json:"author_id" db:"author_id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	Completed  bool      `json:"completed" db:"completed"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	LastUpdate time.Time `json:"last_update" db:"last_update"`
}

type UpdateTodoRequest struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed *bool  `json:"completed"`
}

type InsertTodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64
}
