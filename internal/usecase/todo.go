package usecase

import (
	"github.com/example/go-todo-app/internal/entity"
	"github.com/example/go-todo-app/internal/repository"
)

type TodoUsecase interface {
	Add(title string) (*entity.Todo, error)
	List() ([]*entity.Todo, error)
	Complete(id uint) error
}

type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(r repository.TodoRepository) TodoUsecase {
	return &todoUsecase{repo: r}
}

func (u *todoUsecase) Add(title string) (*entity.Todo, error) {
	todo := &entity.Todo{Title: title}
	if err := u.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) List() ([]*entity.Todo, error) {
	return u.repo.List()
}

func (u *todoUsecase) Complete(id uint) error {
	todos, err := u.repo.List()
	if err != nil {
		return err
	}
	var target *entity.Todo
	for _, t := range todos {
		if t.ID == id {
			target = t
			break
		}
	}
	if target == nil {
		return nil
	}
	target.Completed = true
	return u.repo.Update(target)
}
