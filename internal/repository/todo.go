package repository

import "github.com/example/go-todo-app/internal/entity"

type TodoRepository interface {
	Create(todo *entity.Todo) error
	List() ([]*entity.Todo, error)
	Update(todo *entity.Todo) error
}
