package server

import (
	v1 "AddUser/api/adduser/v1"
	"AddUser/internal/conf"
	"AddUser/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGrpcServer(cf *conf.Config) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(recovery.Recovery()),
	}
	if cf.Address != "" {
		opts = append(opts, grpc.Address(cf.Address))
	}
	if cf.Timeout != nil {
		opts = append(opts, grpc.Timeout(cf.Timeout.AsDuration()))
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterAddUserServer(srv, service.NewAddUser())
	return srv

}
