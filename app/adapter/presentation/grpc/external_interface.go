package grpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/fujimisakari/grpc-todo/app/domain"
)

// Logger is an interface for logging.
type Logger func(ctx context.Context) *zap.Logger

func (f Logger) FromContext(ctx context.Context) *zap.Logger {
	return f(ctx)
}

// Usecase is  an interface for todo usecase.
type Usecase interface {
	GetTodo(ctx context.Context, todoID string) (*domain.Todo, error)
}
