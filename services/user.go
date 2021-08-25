package services

import (
	"context"
	"fmt"

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
