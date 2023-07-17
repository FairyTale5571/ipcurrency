package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fairytale5571/ipcurrency/config"
	"github.com/fairytale5571/ipcurrency/internal/app"
)

// @title IPCurrency
// @version 2.0
// @description API for IPCurrency Service

// @host localhost:3000
// @BasePath /
// @Schemes http

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalf("error while reading config: %s", err.Error())
	}

	application := app.NewApp()
	if err := application.Run(); err != nil {
		log.Fatalf("error while running application: %s", err.Error())
	}

	ctx := registerGracefulHandle()
	<-ctx.Done()

}

func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
