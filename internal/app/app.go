package app

import (
	httpServer "gomarkdownblog/internal/server/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      httpServer.HTTPServer
}

func NewApp() (*App, error) {

	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = newServiceProvider()

	a.serviceProvider.initHTTPRouter()

	return nil
}

func (a *App) initDeps() error {

	err := a.initServiceProvider()
	a.initServer()

	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServer() {

	postRouter := a.serviceProvider.HTTPRouter
	cssRouter := httpServer.NewCSSRouter().Mux

	server := httpServer.NewServer()

	server.AddRouter(postRouter, "")
	server.AddRouter(cssRouter, "/css")

	a.httpServer = *server
}

func (a *App) Run() error {
	if err := a.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
