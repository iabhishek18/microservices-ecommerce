package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":50054")
	grpcServer := grpc.NewServer()
	log.Printf("User service listening on :50054")
	grpcServer.Serve(lis)
}
