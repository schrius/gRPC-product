package main

import (
	"log"
	"net"
	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	//
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
}