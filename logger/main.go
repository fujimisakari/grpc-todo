package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/fujimisakari/grpc-study/logger/pb"
	"github.com/fujimisakari/grpc-study/logger/service"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := &service.LoggerService{}
	pb.RegisterLoggerServer(server, service)
	server.Serve(listenPort)
}
