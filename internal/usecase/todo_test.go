package usecase

import (
	"testing"

	"github.com/example/go-todo-app/internal/entity"
)

type inMemoryRepo struct {
	todos  []*entity.Todo
	nextID uint
}

func newInMemoryRepo() *inMemoryRepo {
	return &inMemoryRepo{nextID: 1}
}

func (r *inMemoryRepo) Create(todo *entity.Todo) error {
	todo.ID = r.nextID
	r.nextID++
	r.todos = append(r.todos, todo)
	return nil
}

func (r *inMemoryRepo) List() ([]*entity.Todo, error) {
	return r.todos, nil
}

func (r *inMemoryRepo) Update(todo *entity.Todo) error {
	for i, t := range r.todos {
		if t.ID == todo.ID {
			r.todos[i] = todo
			return nil
		}
	}
	return nil
}

func TestTodoUsecase_AddAndList(t *testing.T) {
	repo := newInMemoryRepo()
	u := NewTodoUsecase(repo)

	todo, err := u.Add("task1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if todo.ID == 0 || todo.Title != "task1" || todo.Completed {
		t.Errorf("unexpected todo: %+v", todo)
	}

	todos, err := u.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}
	if todos[0].ID != todo.ID || todos[0].Title != "task1" {
		t.Errorf("todo not stored correctly: %+v", todos[0])
	}
}

func TestTodoUsecase_Complete(t *testing.T) {
	repo := newInMemoryRepo()
	u := NewTodoUsecase(repo)

	todo, _ := u.Add("task1")
	_, _ = u.Add("task2")

	if err := u.Complete(todo.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	todos, _ := u.List()
	if !todos[0].Completed {
		t.Errorf("expected todo to be completed")
	}

	if err := u.Complete(999); err != nil {
		t.Fatalf("unexpected error on non-existent id: %v", err)
	}
}
