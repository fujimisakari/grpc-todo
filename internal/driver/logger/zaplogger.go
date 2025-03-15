package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/fujimisakari/grpc-todo/internal/driver/config"
)

func NewZapLogger(env *config.Environment) (*zap.Logger, error) {
	var lv zapcore.Level
	if err := lv.UnmarshalText([]byte(env.LogLevel)); err != nil {
		return nil, err
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(lv)
	if env.IsLocal() {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Development = true
		config.Encoding = "console"
		config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		config.OutputPaths = []string{"stdout"}      // console ouput
		config.ErrorOutputPaths = []string{"stderr"} // console ouput
	}

	return config.Build()
}

type contextKey string

const loggerKey contextKey = "logger"

func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}
