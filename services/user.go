package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/flaviossantana/go-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	//INSERINDO NO BANCO DE DADOS
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "2312",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil

}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserindo",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Usuário foi inserido",
		User: &pb.User{
			Id:    "8970",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Finalizado!",
		User: &pb.User{
			Id:    "8970",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	users := []*pb.User{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Erro ao receber o stream de Usuário:  %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adicionando Usuário: ", req.GetName())

	}

}

func (*UserService) AddUsersStream(stream pb.UserService_AddUsersStreamServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Erro ao receber o stream de User %v", err)

		}

		err = stream.Send(&pb.UserResultStream{
			Status: "201",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Erro ao enviar o stream de User %v", err)

		}
	}
}
