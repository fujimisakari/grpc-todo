package service

import (
	"context"
	"fmt"

	pb "github.com/fujimisakari/grpc-study/logger/pb"
	"github.com/go-redis/redis"
)

const (
	AccessCounter string = "AccessCounter"
)

type LoggerService struct{}

func (s *LoggerService) getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (s *LoggerService) GetCount(ctx context.Context, message *pb.GetCountRequest) (*pb.GetCountResponse, error) {
	c := s.getRedisClient()
	val, err := c.Get(AccessCounter).Result()
	fmt.Println(val)
	if err != nil {
		panic(err)
	}
	return &pb.GetCountResponse{Name: "GetCount", Kind: val}, nil
}

func (s *LoggerService) CountUp(ctx context.Context, message *pb.CountUpMessage) (*pb.CountUpResponse, error) {
	c := s.getRedisClient()
	result, err := c.Incr(AccessCounter).Result()
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return &pb.CountUpResponse{Name: "CountUp", Kind: "result"}, nil
}
