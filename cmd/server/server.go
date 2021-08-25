package main

import (
	"log"
	"net"

	"github.com/flaviossantana/go-grpc/pb"
	"github.com/flaviossantana/go-grpc/services"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Servidor não está onlione!: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro no servidor: %v", err)
	}

}
