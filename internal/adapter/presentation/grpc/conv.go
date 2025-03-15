package grpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/fujimisakari/grpc-todo/internal/adapter/pb"
	"github.com/fujimisakari/grpc-todo/internal/domain"
)

func convTodoPb(todo *domain.Todo) *pb.Todo {
	return &pb.Todo{
		Id:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Priority:    pb.TodoPriority(todo.Priority),
		Completed:   todo.Completed,
		DueTime:     timestamppb.New(todo.DueTime),
		CreatedAt:   timestamppb.New(todo.CreatedAt),
		UpdatedAt:   timestamppb.New(todo.UpdatedAt),
	}
}

func convTodoDomainFromCreateTodoReq(req *pb.CreateTodoRequest) *domain.Todo {
	return &domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    domain.Priority(req.Priority),
		DueTime:     req.DueDate.AsTime(),
	}
}

func convTodoDomainFromUpdateTodoReq(req *pb.UpdateTodoRequest) *domain.Todo {
	return &domain.Todo{
		ID:          req.TodoId,
		Title:       req.Todo.Title,
		Description: req.Todo.Description,
		Priority:    domain.Priority(req.Todo.Priority),
		DueTime:     req.Todo.DueTime.AsTime(),
	}
}
