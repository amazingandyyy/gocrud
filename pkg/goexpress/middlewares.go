package goexpress

import (
	"fmt"
	"net/http"
	"strings"
)

func (e *Express) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s?%s %v\n", r.Method, r.URL.Path, r.URL.RawQuery, r.Proto)
		next.ServeHTTP(w, r)
	})
}

func (e *Express) ParamsParserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, path := range e.paths {
			if r.URL.Path != path.path && strings.Contains(r.URL.Path, path.path) {
				paramsVal := strings.ReplaceAll(r.URL.Path, path.path, "")
				paramsVal = strings.Trim(paramsVal, "/")
				fmt.Println("matched", r.URL.Path, "params", path.paramsKey, paramsVal, path.path)
				path.paramsVal = paramsVal
				http.Redirect(w, r, path.path, http.StatusTemporaryRedirect)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
