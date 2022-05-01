package goexpress

import (
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
	http.Handler
}

func (e *Express) TransformHandler(handler func(*Request, *Response)) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		req := &Request{Request: r}
		if e.paths[r.URL.Path].paramsKey != "" {
			req.SetParam(e.paths[r.URL.Path].paramsKey, e.paths[r.URL.Path].paramsVal)
		}
		res := &Response{w}
		handler(req, res)
	}
	return http.HandlerFunc(fn)
}

func (e *Express) RegisterHandler(method string, path string, handler func(*Request, *Response)) {
	param := ""
	if strings.Contains(path, "/:") {
		param = strings.Split(path, "/:")[1]
		path = strings.Split(path, "/:")[0]
	}
	fmt.Println("RegisterHandler:", method, path, "param:", param)
	p := strings.Split(path, "/:")[0]
	e.paths[path] = &Path{
		path:      path,
		paramsKey: param,
	}
	if _, ok := e.handlers[p]; !ok {
		e.handlers[p] = map[string]http.Handler{}
	}
	e.handlers[p][method] = e.TransformHandler(handler)
}

func transformPath(path string) string {
	return path
}

func (e *Express) Get(path string, handler func(*Request, *Response)) {
	e.RegisterHandler(http.MethodGet, transformPath(path), handler)
}

func (e *Express) Post(path string, handler func(*Request, *Response)) {
	e.RegisterHandler(http.MethodPost, transformPath(path), handler)
}

func (e *Express) Delete(path string, handler func(*Request, *Response)) {
	e.RegisterHandler(http.MethodDelete, transformPath(path), handler)
}

func (e *Express) Put(path string, handler func(*Request, *Response)) {
	e.RegisterHandler(http.MethodPut, transformPath(path), handler)
}
