package internal

import (
	"context"
	"fmt"
	"log"
	"net"

	gpb "github.com/OliverCardoza/test-grpc-http/api/greeting/v0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type greetingService struct {
	gpb.UnimplementedGreetingServiceServer
}

func (g *greetingService) Greeting(ctx context.Context, req *gpb.GreetingRequest) (*gpb.GreetingResponse, error) {
	log.Printf("gRPC: Greeting %s", req.GetName())

	if req.GetName() == "" {
		return nil, fmt.Errorf("error no name in request")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("no metadata :(")
	} else {
		log.Printf("metadata: %v", md)
	}

	return &gpb.GreetingResponse{
		Msg: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}

func RunService(l net.Listener) error {
	s := grpc.NewServer()
	gpb.RegisterGreetingServiceServer(s, &greetingService{})
	log.Printf("gRPC listening")
	err := s.Serve(l)
	return err
}
