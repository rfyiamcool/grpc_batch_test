package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-go-demo/src/helloworld"
)

const (
	defaultName = "world"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:50051", "input server addr")
	clientNum := flag.Int("c", 10, "client num")
	flag.Parse()

	log.Printf("server addr: %s, multi client: %d", *addr, *clientNum)

	oneClient(*addr, *clientNum)
	time.Sleep(30 * time.Second)
	multiClient(*addr, *clientNum)
}

func oneClient(addr string, clientNum int) {
	var wg sync.WaitGroup
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	start := time.Now()
	for index := 0; index < clientNum; index++ {
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

func multiClient(addr string, clientNum int) {
	var wg sync.WaitGroup

	count := clientNum
	clientPool := []*grpc.ClientConn{}
	for index := 0; index < count; index++ {
		clientPool = append(clientPool, newClient(addr))
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

func newClient(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	return conn
}
