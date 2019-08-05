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
./bin/client -addr=172.31.12.12:50051 -c=30 -n=500000 -mode=multi
```

result

```
multi client: 10, qps is 98571
```

## summary

test in the following network environment, the performance is almost the same

* in lan network
* in wan network
* in different city