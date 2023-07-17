package main

import (
	"context"
	"github.com/fairytale5571/ipcurrency/config"
	"github.com/fairytale5571/ipcurrency/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalf("error while reading config: %s", err.Error())
	}

	application := app.NewApp()
	if err := application.Run(); err != nil {
		log.Fatalf("error while running application: %s", err.Error())
	}

	ctx := registerGracefulHandle()
	_ = <-ctx.Done()

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
