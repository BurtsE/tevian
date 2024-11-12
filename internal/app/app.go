package app

import (
	"context"
	"log"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp() *App {
	a := &App{}
	a.serviceProvider = NewSericeProvider()
	return a
}

func (a *App) Run(ctx context.Context) error {
	return a.runServer(ctx)
}

func (a *App) runServer(ctx context.Context) error {
	go a.serviceProvider.TelegramBot().Start(ctx)
	log.Printf("server is running on %s", a.serviceProvider.Config().Server.Host)
	return a.serviceProvider.Router().Start()
}
func (a *App) Stop() error {
	return a.serviceProvider.Router().Shutdown()
}
