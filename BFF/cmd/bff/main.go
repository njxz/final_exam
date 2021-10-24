package bff

import (
	"bff/internal/conf"
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


}
