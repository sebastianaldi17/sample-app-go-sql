package repo

import (
	"github.com/sebastianaldi17/sample-app-go-sql/internal/entity"
)

const (
	queryGetTodos = `
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
	`
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
			id = $1
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

func (r *Repo) GetTodos() ([]entity.Todo, error) {
	todos := make([]entity.Todo, 0)
	err := r.db.Select(&todos, queryGetTodos)
	return todos, err
}

func (r *Repo) GetTodoByID(id int64) (entity.Todo, error) {
	todo := make([]entity.Todo, 0)
	err := r.db.Select(&todo, queryGetTodoByID, id)
	if err != nil {
		return entity.Todo{}, err
	}
	if len(todo) == 0 {
		return entity.Todo{}, nil
	}
	return todo[0], nil
}

func (r *Repo) InsertTodo(req entity.InsertTodoRequest) error {
	_, err := r.db.Exec(queryInsertTodo, req.Title, req.Content, req.UserID)
	return err
}

func (r *Repo) UpdateTodo(req entity.UpdateTodoRequest) error {
	_, err := r.db.Exec(queryUpdateTodo, req.Title, req.Content, req.Completed, req.ID)
	return err
}

func (r *Repo) DeleteTodo(id int64) error {
	_, err := r.db.Exec(queryDeleteTodo, id)
	return err
}

func (r *Repo) GetTodoByAuthor(authorID int64) ([]entity.Todo, error) {
	todo := make([]entity.Todo, 0)
	err := r.db.Select(&todo, queryGetTodoByAuthor, authorID)
	return todo, err
}
