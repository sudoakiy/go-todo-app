package mysql

import (
	"github.com/example/go-todo-app/internal/entity"
	"github.com/example/go-todo-app/internal/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *entity.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) List() ([]*entity.Todo, error) {
	var todos []*entity.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *entity.Todo) error {
	return r.db.Save(todo).Error
}
