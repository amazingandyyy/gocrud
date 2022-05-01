package goexpress

import (
	"net/http"
	"net/url"
)

type Request struct {
	*http.Request
	Params map[string]string
}

func (r *Request) Method() string {
	return r.Request.Method
}

func (r *Request) SetParam(key string, val string) {
	r.Params[key] = val
}

func (r *Request) Query() url.Values {
	return r.URL.Query()
}
