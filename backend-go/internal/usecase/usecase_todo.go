package usecase

import (
	"github.com/redis/go-redis/v9"
	todoEntity "github.com/sebastianaldi17/sample-app-go-sql/internal/entity/todo"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/pkg/logger"
)

func (u *Usecase) GetTodoByID(id int64) (todoEntity.Todo, error) {
	return u.repo.GetTodoByID(id)
}

func (u *Usecase) InsertTodo(req todoEntity.InsertTodoRequest) error {
	err := u.repo.InsertTodo(req)
	if err != nil {
		return err
	}
	u.repo.DeleteTodoByAuthorCache(req.UserID)
	return nil
}

func (u *Usecase) UpdateTodo(req todoEntity.UpdateTodoRequest) error {
	return u.repo.UpdateTodo(req)
}

func (u *Usecase) DeleteTodo(id int64) error {
	return u.repo.DeleteTodo(id)
}

func (u *Usecase) VerifyTodoAuthor(todoID, userID int64) (bool, error) {
	todo, err := u.repo.GetTodoByID(todoID)
	if err != nil {
		return false, err
	}

	return todo.AuthorID == userID, nil
}

func (u *Usecase) GetTodoByAuthor(authorID int64) ([]todoEntity.Todo, error) {
	result := make([]todoEntity.Todo, 0)
	cacheRes, err := u.repo.GetTodoByAuthorCache(authorID)
	if err != nil {
		if err != redis.Nil {
			logger.Error(err.Error())
		}

		fromDb, err := u.repo.GetTodoByAuthor(authorID)
		if err != nil {
			return result, err
		}

		u.repo.SetTodoByAuthorCache(fromDb, authorID)
		result = append(result, fromDb...)
	} else {
		result = append(result, cacheRes...)
	}
	return result, nil
}
