# Go Simple gRPC

## Generate `.proto`
```bash
protoc --go_out=./proto --go-grpc_out=./proto proto/person.proto
```

## Run
### Server
```bash
go run server/main.go
```

### Client
```bash
go run client/main.go
```