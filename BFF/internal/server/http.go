package server

import (
	"bff/api/helloworld/v1"
	"bff/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHttpServer(cf *conf.Config) *http.Server {
	var opt = []http.ServerOption{
		http.Middleware(recovery.Recovery()),
	}
	if cf.Address!=""{
		opt = append(opt, http.Address(cf.Address))
	}
	if cf.Timeout!=nil{
		opt = append(opt, http.Timeout(cf.Timeout.AsDuration()))
	}
	srv:=http.NewServer(opt...)
	return srv
}
