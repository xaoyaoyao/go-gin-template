/**
 * Package server
 * @file      : main.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 09:59
 **/

package main

import (
	"context"
	"fmt"
	"github.com/coverai/api/internal/config"
	"github.com/coverai/api/internal/logs"
	"github.com/coverai/api/internal/service"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	// init config
	err := config.Init()
	if err != nil {
		fmt.Println("Error initializing config:", err)
		return
	}
	logs.Init(config.Get().LogLevel)

	// service start
	if err := service.New(nil).Run(ctx); err != nil {
		panic(err)
	}
}
