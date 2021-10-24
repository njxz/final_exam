package service

import (
	v1 "User/api/user/v1"
	"User/internal/data"
	"context"
)

type User struct {
	v1.UnimplementedUserServer
	data *data.Data
}

func NewUserService(data *data.Data) User {
	user := User{data: data}
	return user
}

func (u User) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	s, err := u.data.SelectName(req.Userid)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserResponse{User: s, Userid: s}, nil
}
func (u User) AddUser(ctx context.Context, req *v1.AddUserRequest) (*v1.AddResponse, error) {
	err := u.data.CreateId(req.Userid, req.User)
	if err != nil {
		return &v1.AddResponse{Status: false}, err
	}
	return &v1.AddResponse{Status: true}, nil
}
