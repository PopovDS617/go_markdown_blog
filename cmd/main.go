package main

import (
	"fmt"
	"gomarkdownblog/internal/logger"
	"gomarkdownblog/internal/middleware"
	"os"

	"log"
	"net/http"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")

	logger := logger.Init()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", middleware.LoggingMiddleware(func(w http.ResponseWriter, r *http.Request) {

	}, logger))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux); err != nil {
		log.Fatal(err)
	}
}
