package goexpress

import (
	"fmt"
	"net/http"
)

type Path struct {
	path      string
	paramsKey string
	paramsVal string
}

type Express struct {
	http     *http.Server
	mux      *http.ServeMux
	paths    map[string]*Path
	handlers map[string]map[string]http.Handler // { path: { method: handler } }
}

func (e *Express) Listen(port int) {
	fmt.Println("Server started on port", port)
	e.http.Addr = fmt.Sprintf(":%d", port)
	//e.http.Handler = e.ParamsParserMiddleware(e.LoggingMiddleware(e.mux))
	e.http.Handler = e.LoggingMiddleware(e.mux)

	for path := range e.paths {
		fmt.Println("path", path)
		e.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			if e.handlers[r.URL.Path][r.Method] != nil {
				e.handlers[r.URL.Path][r.Method].ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
		})
	}

	err := e.http.ListenAndServe()
	if err != nil {
		return
	}
}

func Init() *Express {
	return &Express{
		http:     &http.Server{},
		mux:      http.NewServeMux(),
		paths:    map[string]*Path{},
		handlers: make(map[string]map[string]http.Handler),
	}
}
