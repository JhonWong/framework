package main

import (
	"jwwebframework/framework"
	"jwwebframework/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()
	core.Use(
		middleware.Recovery(),
		middleware.Test1(),
		middleware.Test2())

	registerRoute(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
