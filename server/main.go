package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "grpc_batch_test/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// fmt.Println("######### get client request name :"+in.Name)
	return &pb.HelloReply{Message: "main: Hello " + in.Name}, nil
}

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8081", nil)
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxConcurrentStreams(1024),
		grpc.WriteBufferSize(64*1024),
	)
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
