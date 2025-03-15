package grpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/fujimisakari/grpc-todo/internal/adapter/pb"
	"github.com/fujimisakari/grpc-todo/internal/driver/config"
	"github.com/fujimisakari/grpc-todo/internal/driver/grpc/interceptor"
	"github.com/fujimisakari/grpc-todo/internal/driver/logger"
)

type todoGRPC struct {
	server *grpc.Server
	env    *config.Environment
}

func New(env *config.Environment, todoService pb.TodoServiceServer, log *zap.Logger) *todoGRPC {
	// Add your own interceptor or server option here
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			// set logger to context
			interceptor.WithLoggerInterceptor(log),
			// write request and response logs verbosely with status code
			interceptor.VerboseLoggingUnaryServerInterceptor(),
			interceptor.WithRecoveryInterceptor(),
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterTodoServiceServer(server, todoService)

	return &todoGRPC{env: env, server: server}
}

func (s *todoGRPC) RunServer(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.env.GRPCPort))
		if err != nil {
			return fmt.Errorf("failed to listen on port: %v: %w", s.env.GRPCPort, err)
		}
		logger.FromContext(ctx).Info("gRPC server listening", zap.Int("address", s.env.GRPCPort))
		return s.server.Serve(lis)
	})

	eg.Go(func() error {
		<-ctx.Done()
		s.server.GracefulStop()
		return ctx.Err()
	})

	return eg.Wait()
}
