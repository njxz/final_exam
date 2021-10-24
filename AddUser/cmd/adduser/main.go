package main

import (
	"AddUser/internal/conf"
	"AddUser/internal/server"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"golang.org/x/sync/errgroup"
)

func main() {
	c := config.New(config.WithSource(
		file.NewSource("configs/config.json")))
	var cf conf.Config

	if err := c.Load(); c != nil {
		fmt.Println("config laod error: ", err)
		return
	}
	if err := c.Scan(&cf); err != nil {
		fmt.Println("config scan error: ", err)
		return
	}

	g, err := errgroup.WithContext(context.Background())
	if err != nil {
		fmt.Println("errgroup init error")
		return
	}

	srv := server.NewGrpcServer(&cf)

}
