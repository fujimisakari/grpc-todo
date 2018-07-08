package client

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/fujimisakari/grpc-study/dashboard/pb"
)

func NewDashboard() (pb.DashboardClient, error) {
	conn, err := grpc.Dial("dashboard:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
		return nil, err
	}

	client := pb.NewDashboardClient(conn)
	return client, nil
}
