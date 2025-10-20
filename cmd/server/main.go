package main

import (
	"log"
	"my-service/internal/server"
	pb "my-service/pkg/api/test"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, server.NewServer())

	lis, _ := net.Listen("tcp", ":50051")
	log.Fatal(s.Serve(lis))
}