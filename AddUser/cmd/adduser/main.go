package main

import (
	"AddUser/internal/conf"
	"AddUser/internal/server"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"sync"
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
	ctx := context.Background()
	wg := sync.WaitGroup{}

	srv := server.NewGrpcServer(&cf)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := srv.Start(ctx)
		fmt.Println(err)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		srv.Stop(ctx)
	}()
	wg.Wait()
}
