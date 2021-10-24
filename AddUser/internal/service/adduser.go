package service

import (
	v1 "AddUser/api/adduser/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type AddUser struct {
	v1.UnimplementedAddUserServer
}

func NewAddUser() AddUser {
	return AddUser{}
}

func (a *AddUser) AddUser(ctx context.Context, request *v1.AddRequest) (*v1.AddResponse, error) {
	conn, err := grpc.Dial("127.0.0.1:80")
	if err != nil {
		fmt.Println("grpc client connect error")
		return nil, err
	}
	defer conn.Close()
	client := v1.NewAddUserClient(conn)
	resp, err := client.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp, err
}
