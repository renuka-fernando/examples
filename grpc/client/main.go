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

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	srvStream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
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
		log.Printf("Greeting: %s", r.GetMessage())
	}

	cStream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, name := range []string{"Alice", "Bob", "Charlie"} {
		if err := cStream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	bidiStream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, name := range []string{"Alice", "Bob", "Charlie"} {
		if err := bidiStream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		r, err := bidiStream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		time.Sleep(1 * time.Second)
	}

}
