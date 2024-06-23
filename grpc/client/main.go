package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/renuka-fernando/examples/grpc/server/api/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Renuka"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type client struct {
	ctx    context.Context
	client pb.GreeterClient
}

func (c *client) unaryCall(name string) {
	r, err := c.client.SayHello(c.ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("[unaryCall] Greeting: %s", r.GetMessage())
}

func (c *client) serverStream(name string) {
	srvStream, err := c.client.LotsOfReplies(c.ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		r, err := srvStream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("[serverStream] Greeting: %s", r.GetMessage())
	}
}

func (c *client) clientStream(names []string) {
	stream, err := c.client.LotsOfGreetings(c.ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}

func (c *client) bidiStream(names []string) {
	stream, err := c.client.BidiHello(c.ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		r, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("[bidiStream] Greeting: %s", r.GetMessage())
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	greetClient := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	c := &client{
		ctx:    ctx,
		client: greetClient,
	}

	names := []string{"Alice", "Bob", "Charlie", "four", "five", "six", "seven"}
	c.unaryCall(*name)
	c.unaryCall("World")

	go c.serverStream(*name)
	go c.clientStream(names)
	go c.bidiStream(names)
	time.Sleep(15 * time.Second)
}
