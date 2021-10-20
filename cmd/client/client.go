package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUsersStream(client)

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

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Flávio Santana",
		Email: "flavaio@santana.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Não foi possivel adicionar usuário: %v", err)
	}

	for {
		stream, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Não pode receber a MSG: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())

	}

}

func AddUsers(client pb.UserServiceClient) {

	reqs := usuarios()

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Erro ao criar requisição de Usuário %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		fmt.Println("Enviando a requisição do ", req.Name)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Erro ao receber resposta de Usuário %v", err)
	}

	fmt.Println(res)

}

func AddUsersStream(client pb.UserServiceClient) {

	reqs := usuarios()

	stream, err := client.AddUsersStream(context.Background())

	if err != nil {
		log.Fatalf("Erro ao criar a requisição do Usuário: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Enviando Usuário no Cliente: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Erro ao Receber Usuário do Servidor %v", err)
				break
			}

			fmt.Printf("Recebendo Usuário no Cliente: %v STATUS: %v\n", res.GetUser().GetName(), res.GetStatus())

		}

		close(wait)

	}()

	<-wait

}

func usuarios() []*pb.User {
	return []*pb.User{
		{
			Id:    "1",
			Name:  "Jão Da Silva",
			Email: "jao@email.com",
		},
		{
			Id:    "2",
			Name:  "Maria Da Silva",
			Email: "maria@email.com",
		},
		{
			Id:    "3",
			Name:  "Antoin Da Silva",
			Email: "antoin@email.com",
		},
		{
			Id:    "4",
			Name:  "Pedro Da Silva",
			Email: "pedro@email.com",
		},
		{
			Id:    "5",
			Name:  "Juzé Da Silva",
			Email: "juzeo@email.com",
		},
	}
}
