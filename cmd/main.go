package main

import (
	"gomarkdownblog/internal/app"
	"gomarkdownblog/internal/logger"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		logger.Fatal("failed to init the app " + err.Error())
	}
	if app.Run(); err != nil {
		logger.Fatal("failed to run the app" + err.Error())
	}

}
