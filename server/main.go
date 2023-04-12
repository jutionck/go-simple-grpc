package main

import (
	"context"
	"github.com/jutionck/go-simple-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.PersonRequest) (*proto.PersonResponse, error) {
	person := req.GetPerson()
	message := "Hello " + person.GetName()
	return &proto.PersonResponse{Message: message}, nil
}

func main() {
	// protocol => tcp
	// host & port => : (localhost), port = 1200
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}

	// setup server
	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, &server{})
	log.Println("Server running in localhost:1200")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
