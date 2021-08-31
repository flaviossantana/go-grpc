package main

import (
	"context"
	"log"

	"github.com/flaviossantana/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Não foi possível conectar no servidor gRPC: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)

}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "9999",
		Name:  "João do Gado",
		Email: "jg@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Não foi possível fazer requisição gRPC: %v", err)
	}

	log.Println(res)

}
