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

func (s *HTTPServer) AddRouter(router *http.ServeMux, pattern string) {

	if pattern != "" {

		s.mux.Handle(pattern+"/", http.StripPrefix(pattern, router))
	} else {
		s.mux.Handle("/", router)

	}

}

func (s *HTTPServer) Run() error {
	if err := http.ListenAndServe(":8000", s.mux); err != nil {
		return err
	}
	return nil
}
