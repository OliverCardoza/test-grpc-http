package internal

import (
	"context"
	"fmt"
	"log"
	"net"

	gpb "github.com/OliverCardoza/test-grpc-http/api/greeting/v0"
	"google.golang.org/grpc"
)

type greetingService struct {
	gpb.UnimplementedGreetingServiceServer
}

func (g *greetingService) Greeting(ctx context.Context, req *gpb.GreetingRequest) (*gpb.GreetingResponse, error) {
	log.Printf("gRPC: Greeting %s", req.GetName())

	if req.GetName() == "" {
		return nil, fmt.Errorf("error no name in request")
	}

	return &gpb.GreetingResponse{
		Msg: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}

func RunService(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	gpb.RegisterGreetingServiceServer(s, &greetingService{})
	log.Printf("gRPC listenting on port=%s", port)
	err = s.Serve(lis)
	return err
}
