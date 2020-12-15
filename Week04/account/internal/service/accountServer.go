package service

import (
	"context"
	pb "github.com/netscotte/Go-000/Week04/account/api/account/v1"
	"log"
)

type Server struct {
	*pb.UnimplementedAccountServer
}

func (s *Server) GetUserById(ctx context.Context, id *pb.Id) (*pb.UserInfo, error) {
	log.Printf("client request user for id: %v\n", id)
	// todo: 如何将pb.Id传给pb.UserInfo呢？？
	return &pb.UserInfo{Id: 22, Name: "netliu", Age: 22}, nil
}
