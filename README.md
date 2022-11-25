# test-grpc-http

Test to play around with service gRPC and HTTP off of the same port.

## Usage

Run the server

```bash
go run cmd/server/main.go
```

Make HTTP request

```bash
wget -qO- localhost:12345/hello
```

Make a gRPC request

```bash
grpcurl -plaintext \
    -proto api/greeting/v0/greeting.proto \
    -d '{"name": "World"}' \
    localhost:12346 \
    greeting.GreetingService/Greeting
```

Regenerate protos

```bash
protoc api/greeting/v0/greeting.proto \
    --go_out=. --go-grpc_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative
```

## References

* https://www.cockroachlabs.com/blog/a-tale-of-two-ports/
* https://drgarcia1986.medium.com/listen-grpc-and-http-requests-on-the-same-port-263c40cb45ff
