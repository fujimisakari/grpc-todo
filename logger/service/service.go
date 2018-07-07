package service

import (
	"context"

	pb "github.com/fujimisakari/grpc-study/logger/pb"
)

type LoggerService struct{}

func (s *LoggerService) GetCount(ctx context.Context, message *pb.GetCountRequest) (*pb.GetCountResponse, error) {
	return &pb.GetCountResponse{Name: "GetCount", Kind: "mainecoon"}, nil
}

func (s *LoggerService) CountUp(ctx context.Context, message *pb.CountUpMessage) (*pb.CountUpResponse, error) {
	return &pb.CountUpResponse{Name: "CountUp", Kind: "mainecoon"}, nil
}
