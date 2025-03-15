package interceptor

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/fujimisakari/grpc-todo/internal/driver/logger"
)

func WithRecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err := status.Errorf(codes.Internal, "recovered from panic: %v", r)
				logger.FromContext(ctx).Error("recovered from panic", zap.Any("panic", r), zap.Error(err))
			}
		}()
		return handler(ctx, req)
	}
}
