package main

import (
	"gomarkdownblog/internal/logger"
	"log"
	"net/http"
)

func main() {
	logger := logger.Init()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		logger.Info(r.URL.Path)

	})

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
