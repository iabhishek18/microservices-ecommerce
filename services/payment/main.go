package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50053")
	grpcServer := grpc.NewServer()
	log.Printf("Payment service listening on :50053")
	grpcServer.Serve(lis)
}
