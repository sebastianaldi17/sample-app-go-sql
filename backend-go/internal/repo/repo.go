package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/entity"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

const (
	queryGetTodos = `
		SELECT
			id,
			title,
			content,
			created_at,
			last_update
		FROM
			todos
	`
	queryGetTodoByID = `
		SELECT
			id,
			title,
			content,
			created_at,
			last_update
		FROM
			todos
		WHERE
			id = $1
	`
	queryInsertTodo = `
		INSERT INTO	
			todos(title, content)
		VALUES
			($1, $2)
	`
	queryUpdateTodo = `
		UPDATE
			todos
		SET
			title = COALESCE(NULLIF($1, ''), title),
			content = COALESCE(NULLIF($2, ''), content),
			last_update = now()
		WHERE
			id = $3
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
	_, err := r.db.Exec(queryInsertTodo, req.Title, req.Content)
	return err
}

func (r *Repo) UpdateTodo(req entity.UpdateTodoRequest) error {
	_, err := r.db.Exec(queryUpdateTodo, req.Title, req.Content, req.ID)
	return err
}

func (r *Repo) DeleteTodo(id int64) error {
	_, err := r.db.Exec(queryDeleteTodo, id)
	return err
}
