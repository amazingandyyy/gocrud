package goexpress

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type Response struct {
	http.ResponseWriter
}

func (r *Response) Status(code int) {
	r.WriteHeader(code)
}

func IsJSON(str interface{}) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str.(string)), &js) == nil
}

func (r *Response) Json(data interface{}) (err error) {
	j, err := json.Marshal(data)
	if err != nil {
		return
	}
	r.Header().Set("Content-Type", "application/json")
	_, err = r.Write(j)
	return
}

func (r *Response) Send(data interface{}) (err error) {
	if reflect.TypeOf(data).Kind() == reflect.String {
		if IsJSON(data) {
			r.Header().Set("Content-Type", "application/json")
		} else {
			r.Header().Set("Content-Type", "text/plain")
		}
		_, err = r.Write([]byte(data.(string)))
	}
	return
}
