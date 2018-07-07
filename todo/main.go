package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/fujimisakari/grpc-study/todo/pb"
	"github.com/fujimisakari/grpc-study/todo/service"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := &service.TodoService{}
	pb.RegisterTodoServer(server, service)
	server.Serve(listenPort)
}
