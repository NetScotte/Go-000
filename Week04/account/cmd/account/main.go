package main

import (
	"google.golang.org/grpc"
	pb "github.com/netscotte/Go-000/Week04/account/api/account/v1"
	"net"

	accountServer "github.com/netscotte/Go-000/Week04/account/internal/service"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		panic(err)
	}
	gs := grpc.NewServer()
	pb.RegisterAccountServer(gs, &accountServer.Server{})
	if err := gs.Serve(lis); err != nil {
		panic(err)
	}
}
