package repository

import (
	"context"

	"github.com/fujimisakari/grpc-todo/app/domain"
)

// Repository is an interface for repository.
type Repository interface {
	Todo() Todo
}

type Todo interface {
	// GetTodo gets a todo by id.
	GetTodoByID(ctx context.Context, id string) (*domain.Todo, error)
	// ListTodo lists todos.
	ListTodo(ctx context.Context) ([]*domain.Todo, error)
	// CreateTodo creates a todo.
	CreateTodo(ctx context.Context, todo *domain.Todo) error
	// UpdateTodo updates a todo.
	UpdateTodo(ctx context.Context, todo *domain.Todo) error
	// UpdateTodoStatus updates a todo status.
	UpdateTodoStatus(ctx context.Context, id string, completed bool) error
	// DeleteTodo deletes a todo.
	DeleteTodo(ctx context.Context, id string) error
}
