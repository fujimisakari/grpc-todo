package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/fujimisakari/grpc-todo/internal/adapter/pb"
	"github.com/fujimisakari/grpc-todo/internal/driver/config"
	"github.com/fujimisakari/grpc-todo/internal/driver/logger"
)

type gatewayServer struct {
	server *http.Server
	env    *config.Environment
}

func New(ctx context.Context, env *config.Environment) *gatewayServer {
	// Add your own interceptor or server option here

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// connect gRPC endpoint
	err := pb.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", env.GRPCPort), opts)
	if err != nil {
		logger.FromContext(ctx).Error("failed to register gateway", zap.Error(err))
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", env.HTTPPort),
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	return &gatewayServer{env: env, server: server}
}

func (s *gatewayServer) RunServer(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		logger.FromContext(ctx).Info("starting HTTP/REST gateway server", zap.Int("address", s.env.HTTPPort))
		return s.server.ListenAndServe()
	})

	eg.Go(func() error {
		<-ctx.Done()
		if err := s.server.Shutdown(ctx); err != nil {
			logger.FromContext(ctx).Error("failed to Shutdown", zap.Error(err))
		}
		return ctx.Err()
	})

	return eg.Wait()
}
