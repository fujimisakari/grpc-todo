package usecase

import (
	"context"

	"go.uber.org/zap"
)

// Logger is an interface for logging.
type Logger func(ctx context.Context) *zap.Logger

func (f Logger) FromContext(ctx context.Context) *zap.Logger {
	return f(ctx)
}
