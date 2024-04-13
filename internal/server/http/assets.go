package http

import "net/http"

type CSSRouter struct {
	Mux *http.ServeMux
}

func NewCSSRouter() *CSSRouter {

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("assets/css")))

	return &CSSRouter{
		Mux: mux,
	}

}
