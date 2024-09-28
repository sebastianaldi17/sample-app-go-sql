package repo

import (
	userEntity "github.com/sebastianaldi17/sample-app-go-sql/internal/entity/user"
)

const (
	queryGetPasswordHash = `
		SELECT
			password_hash
		FROM
			users
		WHERE
			username = $1
	`

	queryInsertUser = `
		INSERT INTO users(username, password_hash)
		VALUES ($1, $2)
	`

	queryGetUserIDFromUsername = `
		SELECT
			id
		FROM
			users
		WHERE
			username = $1
	`
)

func (r *Repo) GetPasswordHash(username string) (string, error) {
	var passwordHash string
	err := r.db.Get(&passwordHash, queryGetPasswordHash, username)
	return passwordHash, err
}

func (r *Repo) CreateAccount(req userEntity.Login) error {
	_, err := r.db.Exec(queryInsertUser, req.Username, req.Password)
	return err
}

func (r *Repo) GetUserIDFromUsername(username string) (int64, error) {
	var userID int64
	err := r.db.Get(&userID, queryGetUserIDFromUsername, username)
	return userID, err
}
