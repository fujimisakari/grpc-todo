package interceptor

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	"github.com/fujimisakari/grpc-todo/app/driver/logger"
)

func WithLoggerInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := logger.WithContext(ctx, log)
		return handler(newCtx, req)
	}
}

func VerboseLoggingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)

		f := []zapcore.Field{
			zap.Any("request", req),
			zap.Any("response", resp),
			zap.Error(err),
		}
		logger.FromContext(ctx).Info("verbose log", f...)
		return resp, err
	}
}
