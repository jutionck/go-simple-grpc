package main

import (
	"context"
	"github.com/jutionck/go-simple-grpc/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// grpc.WithInsecure() -> hanya gunakan pada development
	// rekomendasi -> WithTransportCredentials
	conn, err := grpc.Dial("localhost:1200", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect : %v", err)
		return
	}
	defer conn.Close()

	greeterClient := proto.NewGreeterClient(conn)
	personRequest := proto.PersonRequest{Person: &proto.Person{
		Name: "Fadli Rahman",
	}}
	response, err := greeterClient.SayHello(context.Background(), &personRequest)
	if err != nil {
		log.Fatalf("failed to get sayHello : %v", err)
		return
	}

	log.Printf("Greeting: %s\n", response.GetMessage())
}
