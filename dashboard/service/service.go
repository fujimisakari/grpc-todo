package service

import (
	"context"
	"fmt"

	"github.com/fujimisakari/grpc-study/dashboard/client"
	pb "github.com/fujimisakari/grpc-study/dashboard/pb"
	loggerpb "github.com/fujimisakari/grpc-study/logger/pb"
	todopb "github.com/fujimisakari/grpc-study/todo/pb"
)

type DashboardService struct{}

func (s *DashboardService) Get(ctx context.Context, message *pb.GetMessage) (*todopb.GetResponse, error) {
	todoClient, err := client.NewTodo()
	treq := &todopb.GetRequest{TargetCat: "tama"}
	tres, err := todoClient.Get(context.TODO(), treq)
	fmt.Println(tres)

	loggerClient, err := client.NewLogger()
	lreq := &loggerpb.GetCountRequest{TargetCat: "tama"}
	lres, err := loggerClient.GetCount(context.TODO(), lreq)
	fmt.Println(lres)

	res := &todopb.GetResponse{Name: "dashboard", Kind: "mainecoon"}

	return res, err
}
