//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"context"

	"cloud.google.com/go/spanner"

	"github.com/fujimisakari/grpc-todo/internal/domain"
)

// Repository is an interface for repository.
type Repository interface {
	Todo() Todo
}

type Todo interface {
	// GetTodo gets a todo by id.
	FindByID(ctx context.Context, ro YORODB, id string) (*domain.Todo, error)
	// ListTodo lists todos.
	ListTodo(ctx context.Context, ro YORODB) ([]*domain.Todo, error)
	// Insert creates a todo.
	Insert(entity *domain.Todo) *spanner.Mutation
	// Update updates a todo.
	Update(entity *domain.Todo) *spanner.Mutation
	// Deleten deletes a todo.
	Delete(entity *domain.Todo) *spanner.Mutation
}
