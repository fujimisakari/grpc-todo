package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/fujimisakari/grpc-todo/internal/adapter/pb"
)

type todoService struct {
	uc     Usecase
	logger Logger
	pb.UnimplementedTodoServiceServer
}

func NewTodoService(uc Usecase, l Logger) pb.TodoServiceServer {
	return &todoService{
		uc:     uc,
		logger: l,
	}
}

func (s *todoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.TodoResponse, error) {
	todo, err := s.uc.GetTodo(ctx, req.GetTodoId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get todo: %v", err)
	}
	return &pb.TodoResponse{
		Todo: convTodoPb(todo),
	}, nil
}

func (s *todoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	todos, err := s.uc.ListTodo(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list todo: %v", err)
	}
	var pbTodos []*pb.Todo
	for _, todo := range todos {
		pbTodos = append(pbTodos, convTodoPb(todo))
	}
	return &pb.ListTodoResponse{
		Todo: pbTodos,
	}, nil
}

func (s *todoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}

func (s *todoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}

func (s *todoService) UpdateTodotatus(ctx context.Context, req *pb.UpdateTodoStatusRequest) (*pb.TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodotatus not implemented")
}

func (s *todoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
