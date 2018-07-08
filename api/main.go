package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/fujimisakari/grpc-study/api/pb"
	"github.com/fujimisakari/grpc-study/api/service"
)

func serveGRPC(port int) error {
	lp, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	//opts = append(opts, interceptor.WithUnaryServerInterceptors())
	s := grpc.NewServer(opts...)
	apiService := &service.APIService{}
	pb.RegisterAPIServer(s, apiService)
	reflection.Register(s)

	fmt.Printf("Starting gRPC server on %d\n", lp)
	return s.Serve(lp)
}

func serveHTTP(grpcPort, httpPort int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	muxOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	})
	mux := runtime.NewServeMux(muxOptions)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("127.0.0.1:%d", grpcPort)
	if err := pb.RegisterAPIHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	fmt.Printf("Start HttpServer Listening on %v\n", httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}

func main() {
	port := 8000
	grpcPort := 8080

	errorch := make(chan error)
	go func() {
		errorch <- serveGRPC(grpcPort)
	}()

	go func() {
		errorch <- serveHTTP(grpcPort, port)
	}()

	for err := range errorch {
		log.Fatal(err)
	}
}
