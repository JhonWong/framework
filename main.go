package main

import (
	"context"
	"jwwebframework/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRoute(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	go func() {
		server.ListenAndServe()
	}()

	//监听系统关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
