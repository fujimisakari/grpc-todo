package service

import (
	"context"
	"fmt"

	loggerpb "github.com/fujimisakari/grpc-study/logger/pb"
	"github.com/fujimisakari/grpc-study/todo/client"
	pb "github.com/fujimisakari/grpc-study/todo/pb"
)

type TodoService struct{}

func (s *TodoService) Get(ctx context.Context, message *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Name: "Get", Kind: "mainecoon"}, nil
}

func (s *TodoService) Create(ctx context.Context, message *pb.CreateRequest) (*pb.CreateResponse, error) {
	c, err := client.NewLogger()
	req := &loggerpb.CountUpMessage{TargetCat: "tama"}
	res, err := c.CountUp(context.TODO(), req)
	fmt.Println(res)
	return &pb.CreateResponse{Name: "Create", Kind: "mainecoon"}, err
}

func (s *TodoService) Delete(ctx context.Context, message *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{Name: "Delete", Kind: "mainecoon"}, nil
}
