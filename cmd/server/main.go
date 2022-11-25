package main

import (
	"log"
	"net/http"

	"github.com/OliverCardoza/test-grpc-http/internal"
	"golang.org/x/sync/errgroup"
)

const (
	httpPort = "12345"
	grpcPort = "12346"
)

func main() {
	g := &errgroup.Group{}
	g.Go(startHTTP)
	g.Go(startGRPC)

	err := g.Wait()
	log.Printf("End error: %v", err)
}

func startHTTP() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP: %v %v", r.Method, r.URL.String())
	})

	log.Printf("HTTP listening on port=%s", httpPort)
	return http.ListenAndServe(":"+httpPort, nil)
}

func startGRPC() error {
	return internal.RunService(grpcPort)
}
