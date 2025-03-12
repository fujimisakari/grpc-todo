package spanner

import (
	"context"

	"github.com/fujimisakari/grpc-todo/app/domain"
	"github.com/fujimisakari/grpc-todo/app/domain/repository"
)

// Repository is an interface for repository.
type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Todo() repository.Todo {
	return &Todo{}
}

type Todo struct{}

// GetTodo gets a todo by id.
func (t *Todo) GetTodoByID(ctx context.Context, id string) (*domain.Todo, error) {
	return nil, nil
}

// ListTodo lists todos.
func (t *Todo) ListTodo(ctx context.Context) ([]*domain.Todo, error) {
	return nil, nil
}

// CreateTodo creates a todo.
func (t *Todo) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	return nil
}

// UpdateTodo updates a todo.
func (t *Todo) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return nil
}

// UpdateTodoStatus updates a todo status.
func (t *Todo) UpdateTodoStatus(ctx context.Context, id string, completed bool) error {
	return nil
}

// DeleteTodo deletes a todo.
func (t *Todo) DeleteTodo(ctx context.Context, id string) error {
	return nil
}
