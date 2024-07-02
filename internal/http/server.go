package http

import (
	"fmt"
	"net/http"
	"os"
)

type HTTPServer struct {
	server *http.Server
	mux    *http.ServeMux
	port   string
}

func NewServer() *HTTPServer {

	httpPort := os.Getenv("HTTP_PORT")

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: mux,
	}

	return &HTTPServer{
		server: server,
		mux:    mux,
		port:   httpPort,
	}

}

func (s *HTTPServer) AddRouter(router http.Handler, pattern string) {

	if pattern != "" {
		s.mux.Handle(pattern+"/", http.StripPrefix(pattern, router))
	} else {
		s.mux.Handle("/", router)

	}

}

func (s *HTTPServer) Run() error {

	return s.server.ListenAndServe()
}
