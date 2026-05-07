package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

type OrderServer struct{}

type SagaOrchestrator struct {
	Steps []SagaStep
}

type SagaStep struct {
	Name     string
	Execute  func() error
	Rollback func() error
}

func (s *SagaOrchestrator) Run() error {
	var completed []int
	for i, step := range s.Steps {
		if err := step.Execute(); err != nil {
			log.Printf("Step %s failed: %v, rolling back...", step.Name, err)
			for j := len(completed) - 1; j >= 0; j-- {
				s.Steps[completed[j]].Rollback()
			}
			return err
		}
		completed = append(completed, i)
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50052")
	grpcServer := grpc.NewServer()
	log.Printf("Order service listening on :50052")
	grpcServer.Serve(lis)
}
