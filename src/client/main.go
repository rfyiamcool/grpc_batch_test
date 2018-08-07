package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-go-demo/src/helloworld"
	"log"
	"sync"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	oneClient()
	time.Sleep(3 * time.Second)
	multiClient()
}

func oneClient() {
	var wg sync.WaitGroup
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	start := time.Now()
	for index := 0; index < 5; index++ {
		var err error
		var r *pb.HelloReply
		go func() {
			for idx := 0; idx < 20000; idx++ {
				r, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				if r.Message != "main: Hello world" {
					log.Printf("####### get server Greeting response: %s", r.Message)
				}
			}
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()
	end := time.Since(start)
	log.Println("one clinet took: ", end)
}

func multiClient() {
	var wg sync.WaitGroup

	count := 5
	clientPool := []*grpc.ClientConn{}
	for index := 0; index < count; index++ {
		clientPool = append(clientPool, newClient())
	}

	start := time.Now()
	for index := 0; index < count; index++ {
		go func(index int) {
			var err error
			var r *pb.HelloReply

			conn := clientPool[index]
			c := pb.NewGreeterClient(conn)

			for idx := 0; idx < 20000; idx++ {
				r, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				if r.Message != "main: Hello world" {
					log.Printf("####### get server Greeting response: %s", r.Message)
				}
			}

			wg.Done()
		}(index)
		wg.Add(1)
	}
	wg.Wait()
	end := time.Since(start)
	log.Println("multi client took: ", end)
}

func newClient() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	return conn
}
