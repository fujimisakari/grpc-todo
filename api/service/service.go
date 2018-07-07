package service

import (
	"context"
	"fmt"

	"github.com/fujimisakari/grpc-study/api/client"
	pb "github.com/fujimisakari/grpc-study/api/pb"
	dpb "github.com/fujimisakari/grpc-study/dashboard/pb"
)

type APIService struct{}

func (s *APIService) GetDashboard(ctx context.Context, message *pb.EmptyRequest) (*pb.DashboardResponse, error) {
	c, err := client.NewDashboard()
	m := &dpb.GetMessage{TargetCat: "aa"}
	res, err := c.Get(ctx, m)
	fmt.Println(res)
	return &pb.DashboardResponse{Status: true}, err
}

func (s *APIService) CreateTodo(ctx context.Context, message *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	return &pb.CreateTodoResponse{Status: true}, nil
}

func (s *APIService) DeleteTodo(ctx context.Context, message *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	return &pb.DeleteTodoResponse{Status: true}, nil
}
