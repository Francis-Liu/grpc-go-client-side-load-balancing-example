package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Francis-Liu/grpc-go-client-side-load-balancing-example/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcPort   = flag.Int("grpcport", 50051, "Server listening port")
	serverName = flag.String("servername", "server1", "Server host name")
)

func init() {
	flag.Parse()
}

type server struct {
}

func (s *server) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: "Hello " + in.GetMessage() + " from " + *serverName}, nil
}

func serve(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	fmt.Printf("Starting server on %d ...\n", *grpcPort)
	serve(*grpcPort)
}
