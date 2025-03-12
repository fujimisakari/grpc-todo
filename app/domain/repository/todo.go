package repository

import (
	"context"
)

// Repository is an interface for repository.
type Todo interface {
	// GetTodo gets a todo by id.
	GetTodo(ctx context.Context, id string) (*Todo, error)
	// ListTodo lists todos.
	ListTodo(ctx context.Context) ([]*Todo, error)
	// CreateTodo creates a todo.
	CreateTodo(ctx context.Context, todo *Todo) (*Todo, error)
	// UpdateTodo updates a todo.
	UpdateTodo(ctx context.Context, todo *Todo) (*Todo, error)
	// UpdateTodoStatus updates a todo status.
	UpdateTodoStatus(ctx context.Context, id string, completed bool) (*Todo, error)
	// DeleteTodo deletes a todo.
	DeleteTodo(ctx context.Context, id string) error
}
