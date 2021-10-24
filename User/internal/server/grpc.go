package server

import (
	v1 "User/api/user/v1"
	"User/internal/conf"
	"User/internal/data"
	"User/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGrpcServer(c *conf.Config, data *data.Data) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery()),
	}
	if c.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Timeout.AsDuration()))
	}
	if c.Address != "" {
		opts = append(opts, grpc.Address(c.Address))
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, service.NewUserService(data))
	return srv

}
