package client

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/fujimisakari/grpc-study/todo/pb"
)

func NewTodo() (pb.TodoClient, error) {
	conn, err := grpc.Dial("todo:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
		return nil, err
	}

	client := pb.NewTodoClient(conn)
	return client, nil
}
