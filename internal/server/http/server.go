package http

import (
	"net/http"
)

type HTTPServer struct {
	mux *http.ServeMux
}

func NewServer(router *http.ServeMux) *HTTPServer {

	return &HTTPServer{
		mux: router,
	}

}

func (s *HTTPServer) Run() error {
	if err := http.ListenAndServe(":8000", s.mux); err != nil {
		return err
	}
	return nil
}
