package main

import (
	"log"
	"net"
	"net/http"

	"github.com/OliverCardoza/test-grpc-http/internal"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
)

const (
	port = "12345"
)

func main() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	// Doesn't work, maybe https://github.com/soheilhy/cmux/issues/95
	// grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcListener := m.Match(cmux.HTTP2())
	anyListener := m.Match(cmux.Any())

	g := &errgroup.Group{}
	g.Go(func() error { return startGRPC(grpcListener) })
	g.Go(func() error { return startHTTP(anyListener) })
	g.Go(func() error {
		log.Printf("TCP listening on port=%s", port)
		return m.Serve()
	})

	err = g.Wait()
	log.Printf("End error: %v", err)
}

func startHTTP(l net.Listener) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP: %v %v", r.Method, r.URL.String())
	})

	log.Printf("HTTP listening")
	return http.Serve(l, nil)
}

func startGRPC(l net.Listener) error {
	return internal.RunService(l)
}
