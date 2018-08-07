# grpc batch test

batch grpc server by one or multi client connnet, then batch test in different quality network.

## get

```
cd $GOPATH;git clone https://github.com/rfyiamcool/grpc_batch_test.git
```

## make go proto

```
cd src; make pb; ls helloworld
```

## test

start server

```
go run src/server/main.go
```

start client

```
go run src/client/main.go
```

result

```
2018/08/07 22:29:16 one clinet took:  6.015762269s
2018/08/07 22:29:25 multi client took:  6.423514978s
```

## summary

none