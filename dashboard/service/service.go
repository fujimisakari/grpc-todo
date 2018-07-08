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

func (s *DashboardService) Get(ctx context.Context, message *pb.GetMessage) (*pb.Response, error) {
	todoClient, err := client.NewTodo()
	treq := &todopb.GetRequest{TargetCat: "tama"}
	tres, err := todoClient.Get(ctx, treq)
	fmt.Println(tres)

	loggerClient, err := client.NewLogger()

	lrem := &loggerpb.CountUpMessage{TargetCat: "tama"}
	h, err := loggerClient.CountUp(ctx, lrem)
	fmt.Println(h)

	lreq := &loggerpb.GetCountRequest{TargetCat: "tama"}
	lres, err := loggerClient.GetCount(ctx, lreq)
	fmt.Println(lres)

	res := &pb.Response{Name: "dashboard", Kind: "mainecoon"}

	return res, err
}
