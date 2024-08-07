package usecase

import (
	"github.com/sebastianaldi17/sample-app-go-sql/internal/entity"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/repo"
)

type Usecase struct {
	repo repo.Repo
}

func New(repo repo.Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetTodos() ([]entity.Todo, error) {
	return u.repo.GetTodos()
}

func (u *Usecase) GetTodoByID(id int64) (entity.Todo, error) {
	return u.repo.GetTodoByID(id)
}

func (u *Usecase) InsertTodo(req entity.InsertTodoRequest) error {
	return u.repo.InsertTodo(req)
}

func (u *Usecase) UpdateTodo(req entity.UpdateTodoRequest) error {
	return u.repo.UpdateTodo(req)
}

func (u *Usecase) DeleteTodo(id int64) error {
	return u.repo.DeleteTodo(id)
}
