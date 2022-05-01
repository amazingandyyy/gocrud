package main

import (
	express "gocrud/pkg/goexpress"
)

func health(req *express.Request, res *express.Response) {
	err := res.Json(map[string]string{"status": "ok"})
	if err != nil {
		return
	}
}

//func hello(req *express.Request, res *express.Response) {
//	err := res.Json(map[string]string{"person": req.Params["person"], "age": req.Query()["age"][0]})
//	if err != nil {
//		return
//	}
//}

func main() {
	app := express.Init()
	//app.Get("/hello/:id", hello)
	app.Post("/health", health)
	app.Get("/health", health)
	app.Listen(8080)
}
