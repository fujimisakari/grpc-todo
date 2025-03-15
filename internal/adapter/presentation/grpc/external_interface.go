package grpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/fujimisakari/grpc-todo/internal/domain"
)

// Logger is an interface for logging.
type Logger func(ctx context.Context) *zap.Logger

func (f Logger) FromContext(ctx context.Context) *zap.Logger {
	return f(ctx)
}

// Usecase is  an interface for todo usecase.
type Usecase interface {
	GetTodo(ctx context.Context, todoID string) (*domain.Todo, error)
	ListTodo(ctx context.Context) ([]*domain.Todo, error)
	UpdateTodo(ctx context.Context, todo *domain.Todo) error
	UpdateTodoStatus(ctx context.Context, todoID string, completed bool) error
	DeleteTodo(ctx context.Context, todoID string) error
}
