package grpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/fujimisakari/grpc-todo/app/adapter/pb"
	"github.com/fujimisakari/grpc-todo/app/domain"
)

func convTodoPb(todo *domain.Todo) *pb.Todo {
	return &pb.Todo{
		TodoId:      todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Priority:    pb.TodoPriority(todo.Priority),
		Completed:   todo.Completed,
		DueDate:     timestamppb.New(todo.DueTime),
		CreatedAt:   timestamppb.New(todo.CreatedAt),
	}
}
