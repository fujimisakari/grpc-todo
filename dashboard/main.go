package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/fujimisakari/grpc-study/dashboard/pb"
	"github.com/fujimisakari/grpc-study/dashboard/service"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := &service.DashboardService{}
	pb.RegisterDashboardServer(server, service)
	server.Serve(listenPort)
}
