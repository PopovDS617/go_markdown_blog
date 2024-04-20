package http

import (
	"net/http"
)

type HTTPServer struct {
	mux *http.ServeMux
}

func NewServer() *HTTPServer {

	mux := http.NewServeMux()

	return &HTTPServer{
		mux,
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

	server := http.Server{
		Addr:    ":8000",
		Handler: s.mux,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
