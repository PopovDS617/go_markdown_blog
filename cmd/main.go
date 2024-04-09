package main

import (
	"gomarkdownblog/internal/app"
	"gomarkdownblog/internal/logger"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		logger.Fatal("err")
	}
	if app.Run(); err != nil {
		logger.Fatal("err")
	}

}
