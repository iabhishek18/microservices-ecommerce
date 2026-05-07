package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

type ProductServer struct{}

func (s *ProductServer) GetProduct(ctx context.Context, req *GetProductRequest) (*Product, error) {
	return &Product{Id: req.Id, Name: "Sample Product", Price: 29.99, Stock: 100}, nil
}

func (s *ProductServer) ListProducts(ctx context.Context, req *ListProductsRequest) (*ProductList, error) {
	products := []*Product{
		{Id: "1", Name: "Laptop", Price: 999.99, Stock: 50, Category: "Electronics"},
		{Id: "2", Name: "Headphones", Price: 79.99, Stock: 200, Category: "Electronics"},
	}
	return &ProductList{Products: products, Total: int32(len(products))}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil { log.Fatalf("failed to listen: %v", err) }
	grpcServer := grpc.NewServer()
	log.Printf("Product service listening on :50051")
	if err := grpcServer.Serve(lis); err != nil { log.Fatalf("failed to serve: %v", err) }
}
