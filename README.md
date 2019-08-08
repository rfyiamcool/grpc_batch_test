# grpc batch test

batch grpc server by one or multi client connnet, then batch test in different quality network.

## get

```
cd $GOPATH/src;git clone https://github.com/rfyiamcool/grpc_batch_test.git
```

## make go proto

```
cd grpc_batch_test; make pb; make build; ls helloworld
```

## test

start server

```
./bin/server
```

start client

```
./bin/client  -h
Usage of ./client:
  -addr string
        input server addr (default "127.0.0.1:50051")
  -c int
        grpc-client connnect num (default 10)
  -g int
        goroutine nums (default 10)
  -n int
        total requests (default 200000)

./bin/client -addr=172.31.12.12:50051 -c=50 -n=1000000 -g=200
```

## result

./bin/client  -addr=10.2.0.24:50051 -c=50 -n=1000000 -g=100

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 50, worker num: 100
multi client: 50, qps is 144601
```

./bin/client  -addr=10.2.0.24:50051 -c=50 -n=1000000 -g=150

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 50, worker num: 150
multi client: 50, qps is 174017
```

./client  -addr=10.2.0.24:50051 -c=50 -n=1000000 -g=400

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 50, worker num: 400
multi client: 50, qps is 228953
```

./client  -addr=10.2.0.24:50051 -c=100 -n=1000000 -g=400

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 100, worker num: 400
multi client: 100, qps is 287921
```

./client  -addr=10.2.0.24:50051 -c=100 -n=1000000 -g=500

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 100, worker num: 500
multi client: 100, qps is 335508
```

./client  -addr=10.2.0.24:50051 -c=150 -n=1000000 -g=1000

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 150, worker num: 600
multi client: 150, qps is 374343
```

./client  -addr=10.2.0.24:50051 -c=200 -n=1000000 -g=1500

```
server addr: 10.2.0.24:50051, totalCount: 1000000, multi client: 200, worker num: 1500
multi client: 200, qps is 394520
```

## summary

test in the following network environment, the performance is almost the same

* in lan network
* in wan network
* in different city