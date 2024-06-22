## gRPC Example

### Generate pb files

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    server/api/helloworld/helloworld.proto
```

### Run Server

```sh
cd server
go run .
```

### Run Client

```sh
cd client
go run .
```
