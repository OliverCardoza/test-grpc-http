package main

import (
	"context"
	"log"
	"time"

	gpb "github.com/OliverCardoza/test-grpc-http/api/greeting/v0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = "12345"
)

func main() {
	ctx := context.Background()

	// Create gRPC connection
	conn, err := grpc.DialContext(
		ctx,
		":"+port,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(intercept),
		grpc.WithTimeout(time.Second),
	)

	if err != nil {
		log.Fatalf("error dialing: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := gpb.NewGreetingServiceClient(conn)

	resp, err := client.Greeting(ctx, &gpb.GreetingRequest{
		Name: "Go Client",
	})
	if err != nil {
		log.Fatalf("grpc error: %v", err)
	}
	log.Printf("greeing response: %v", resp)
}

func intercept(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, _ := metadata.FromOutgoingContext(ctx)
	log.Printf("method=%q, md=%v", method, md)
	return invoker(ctx, method, req, reply, cc, opts...)
}
