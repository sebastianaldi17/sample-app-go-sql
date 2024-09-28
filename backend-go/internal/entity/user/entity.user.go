package user

import "time"

type User struct {
	ID           int64     `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	LastLogin    time.Time `json:"last_login" db:"last_login"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
