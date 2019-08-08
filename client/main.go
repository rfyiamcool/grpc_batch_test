package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "grpc_batch_test/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr       = flag.String("addr", "127.0.0.1:50051", "input server addr")
	clientNum  = flag.Int("c", 10, "grpc-client connnect num")
	workerNum  = flag.Int("g", 10, "goroutine nums")
	totalCount = flag.Int("n", 200000, "total requests")
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	flag.Parse()
	log.Printf("server addr: %s, totalCount: %d, multi client: %d, worker num: %d",
		*addr,
		*totalCount,
		*clientNum,
		*workerNum,
	)

	startTime := time.Now()
	handleMultiClient(*addr, *totalCount, *clientNum, *workerNum)
	costTime := float64(time.Now().Sub(startTime).Nanoseconds()) / float64(1000) / float64(1000) / float64(1000)

	qps := float64(*totalCount) / costTime
	log.Printf("multi client: %d, qps is %.0f", *clientNum, qps)
}

func handleMultiClient(addr string, totalCount, clientNum, workerNum int) {
	var wg sync.WaitGroup

	clientPool := []*grpc.ClientConn{}
	for index := 0; index < clientNum; index++ {
		clientPool = append(clientPool, newClient(addr))
	}

	for index := 0; index < workerNum; index++ {
		wg.Add(1)
		go func(index int) {
			var r *pb.HelloReply
			var err error

			cidx := index % clientNum
			c := pb.NewGreeterClient(clientPool[cidx])

			roundCount := totalCount / workerNum
			for idx := 0; idx < roundCount; idx++ {
				r, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}

				if r.Message != "main: Hello world" {
					log.Printf("get server Greeting response: %s\n", r.Message)
				}
			}

			wg.Done()
		}(index)
	}
	wg.Wait()
}

func newClient(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	return conn
}
