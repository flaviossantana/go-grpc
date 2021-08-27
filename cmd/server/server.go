package main

import (
	"log"
	"net"

	"github.com/flaviossantana/go-grpc/pb"
	"github.com/flaviossantana/go-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log.Printf("Toca pra MUE!")

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Servidor não está onlione!: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro no servidor: %v", err)
	}

}
