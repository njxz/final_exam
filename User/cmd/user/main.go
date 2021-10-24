package main

import (
	"User/internal/conf"
	data2 "User/internal/data"
	"User/internal/server"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"golang.org/x/sync/errgroup"
)

func main() {
	c := config.New(
		config.WithSource(
			file.NewSource("configs/config.json"),
		),
	)
	ctx, _ := context.WithCancel(context.Background())
	var cf conf.Config
	if err := c.Load(); err != nil {
		fmt.Println("fileerr:", err)
		return
	}
	if err := c.Scan(&cf); err != nil {
		fmt.Println(err)
		return
	}
	g, ctx := errgroup.WithContext(ctx)
	data, cleanup, err := data2.NewData(&cf)
	if err != nil {
		fmt.Println("database error: ", err)
		return
	}
	defer cleanup()
	srv := server.NewGrpcServer(&cf, data)

	g.Go(func() error {
		return srv.Start(ctx)
	})
	g.Go(func() error {
		<-ctx.Done()
		return srv.Stop(ctx)
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("quit success")
	}

}
