package client

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/fujimisakari/grpc-study/logger/pb"
)

func NewLogger() (pb.LoggerClient, error) {
	conn, err := grpc.Dial("logger:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
		return nil, err
	}

	client := pb.NewLoggerClient(conn)
	return client, nil
}
