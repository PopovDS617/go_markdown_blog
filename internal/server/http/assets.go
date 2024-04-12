package http

import "net/http"

type Router struct {
	Mux *http.ServeMux
}

func NewCSSRouter() *Router {

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("assets/css")))

	return &Router{
		Mux: mux,
	}

}

func NewImageRouter() *Router {

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("assets/images")))

	return &Router{
		Mux: mux,
	}

}
