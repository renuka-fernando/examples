package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/renuka-fernando/examples/grpc/server/api/helloworld"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func (s *server) LotsOfReplies(req *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.HelloResponse{Message: fmt.Sprintf("%d. Hello %s!", i, req.GetName())}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}
		log.Printf("Received: %v", req.GetName())
	}
}

func (s *server) BidiHello(stream pb.Greeter_BidiHelloServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}
		log.Printf("Received: %v", req.GetName())
		if err := stream.Send(&pb.HelloResponse{Message: "Hello " + req.GetName()}); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
