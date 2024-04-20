package app

import (
	"gomarkdownblog/internal/middleware"
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
	imagesRouter := httpServer.NewImageRouter().Mux

	middlewareChain := middleware.CreateMiddlewareStack(middleware.Caching, middleware.CORS)

	server := httpServer.NewServer()

	server.AddRouter(middlewareChain(postRouter), "")
	server.AddRouter(middlewareChain(cssRouter), "/css")
	server.AddRouter(middlewareChain(imagesRouter), "/images")

	a.httpServer = *server
}

func (a *App) Run() error {
	if err := a.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
