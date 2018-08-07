# grpc batch test

batch grpc server by one client connect and multi connnet, then batch in different quality network.

## make go proto

```
make pb
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