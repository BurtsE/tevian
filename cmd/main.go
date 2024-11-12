package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tevian/internal/app"
)

func main() {
	a := app.NewApp()
	if a == nil {
		log.Fatalf("failed to create app")
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return a.Run(ctx)
	})
	g.Go(func() error {
		<-gCtx.Done()
		return a.Stop()
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
