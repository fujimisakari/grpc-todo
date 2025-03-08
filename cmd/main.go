package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	grpc_service "github.com/fujimisakari/grpc-todo/app/adapter/presentation/grpc"
	"github.com/fujimisakari/grpc-todo/app/driver/config"
	"github.com/fujimisakari/grpc-todo/app/driver/grpc"
	"github.com/fujimisakari/grpc-todo/app/driver/http"
	"github.com/fujimisakari/grpc-todo/app/driver/logger"
	"github.com/fujimisakari/grpc-todo/app/usecase"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

func main() {
	// Create separate main instead of doing the actual code here
	// since os.Exit can not handle `defer`. DON'T call `os.Exit` in the any other place.
	os.Exit(realMain(os.Args))
}

func realMain(_ []string) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Read configurations from environmental variables.
	env, err := config.ReadFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to read env vars: %s\n", err)
		return exitError
	}

	// Setup new zap logger. This logger should be used for all logging in this service.
	// The log level can be updated via environment variables.
	log, err := logger.NewZapLogger(env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup logger: %s\n", err)
		return exitError
	}
	defer func() {
		_ = log.Sync()
	}()
	ctx = logger.WithContext(ctx, log)

	logger.FromContext(ctx).Info("starting server...")

	// Add your own profiler or clinet or worker here.

	// Create new GRPC server
	usecase := usecase.NewUsecase(logger.FromContext)
	todoService := grpc_service.NewTodoService(usecase, logger.FromContext)
	todoGRPC := grpc.New(env, todoService, log)

	// Create new Gateway server
	gateway := http.New(ctx, env)

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error { return todoGRPC.RunServer(ctx) })
	wg.Go(func() error { return gateway.RunServer(ctx) })

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	sigtermFn := func() { cancel() }
	go signalHandler(sigCh, sigtermFn)

	if err := wg.Wait(); err != nil {
		logger.FromContext(ctx).Error("unhandled error received", zap.Error(err))
		return exitError
	}

	return exitOK
}

func signalHandler(ch <-chan os.Signal, sigtermFn func()) {
	for sig := range ch {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, os.Interrupt:
			sigtermFn()
		}
	}
}
