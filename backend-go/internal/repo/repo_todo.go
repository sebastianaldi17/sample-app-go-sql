package repo

import (
	"context"
	"encoding/json"
	"fmt"

	todoEntity "github.com/sebastianaldi17/sample-app-go-sql/internal/entity/todo"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/pkg/logger"
)

const (
	queryGetTodoByID = `
		SELECT
			id,
			author_id,
			title,
			content,
			completed,
			created_at,
			last_update
		FROM
			todos
		WHERE
			id = $1
	`

	queryGetTodoByAuthor = `
		SELECT
			id,
			author_id,
			title,
			content,
			completed,
			created_at,
			last_update
		FROM
			todos
		WHERE
			author_id = $1
		ORDER BY id
	`
	queryInsertTodo = `
		INSERT INTO	
			todos(title, content, author_id)
		VALUES
			($1, $2, $3)
	`
	queryUpdateTodo = `
		UPDATE
			todos
		SET
			title = COALESCE(NULLIF($1, ''), title),
			content = COALESCE(NULLIF($2, ''), content),
			completed = COALESCE($3, completed),
			last_update = now()
		WHERE
			id = $4
	`
	queryDeleteTodo = `
		DELETE FROM
			todos
		WHERE
			id = $1
	`
)

func (r *Repo) GetTodoByID(id int64) (todoEntity.Todo, error) {
	todo := make([]todoEntity.Todo, 0)
	err := r.db.Select(&todo, queryGetTodoByID, id)
	if err != nil {
		return todoEntity.Todo{}, err
	}
	if len(todo) == 0 {
		return todoEntity.Todo{}, nil
	}
	return todo[0], nil
}

func (r *Repo) InsertTodo(req todoEntity.InsertTodoRequest) error {
	_, err := r.db.Exec(queryInsertTodo, req.Title, req.Content, req.UserID)
	return err
}

func (r *Repo) UpdateTodo(req todoEntity.UpdateTodoRequest) error {
	_, err := r.db.Exec(queryUpdateTodo, req.Title, req.Content, req.Completed, req.ID)
	return err
}

func (r *Repo) DeleteTodo(id int64) error {
	_, err := r.db.Exec(queryDeleteTodo, id)
	return err
}

func (r *Repo) GetTodoByAuthor(authorID int64) ([]todoEntity.Todo, error) {
	todo := make([]todoEntity.Todo, 0)
	err := r.db.Select(&todo, queryGetTodoByAuthor, authorID)
	return todo, err
}

func (r *Repo) SetTodoByAuthorCache(todos []todoEntity.Todo, authorID int64) {
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = r.redis.SetEx(context.Background(), fmt.Sprintf(todoEntity.TodoByAuthorKeyFmt, authorID), string(jsonBytes), r.defaultTTL).Err()
	if err != nil {
		logger.Error(err.Error())
	}
}

func (r *Repo) GetTodoByAuthorCache(authorID int64) ([]todoEntity.Todo, error) {
	jsonStr, err := r.redis.Get(context.Background(), fmt.Sprintf(todoEntity.TodoByAuthorKeyFmt, authorID)).Result()
	if err != nil {
		return []todoEntity.Todo{}, err
	}

	results := make([]todoEntity.Todo, 0)
	err = json.Unmarshal([]byte(jsonStr), &results)
	return results, err
}

func (r *Repo) DeleteTodoByAuthorCache(authorID int64) {
	err := r.redis.Del(context.Background(), fmt.Sprintf(todoEntity.TodoByAuthorKeyFmt, authorID)).Err()
	if err != nil {
		logger.Error(err.Error())
	}
}
