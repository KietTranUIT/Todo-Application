package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/conf"
)

type OperateHTTPServer interface {
	Start()
	Stop()
}

type HTTPServer struct {
	server http.Server
}

func NewHTTPServer(conf conf.ConfigServer, mux *http.ServeMux) OperateHTTPServer {
	return &HTTPServer{
		server: http.Server{
			Addr:    conf.Addr(),
			Handler: mux,
		},
	}
}

func (srv *HTTPServer) Start() {
	go srv.Stop()
	if err := srv.server.ListenAndServe(); err != nil {
		log.Println("Server error: ", err.Error())
	}
}

func (srv *HTTPServer) Stop() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shut down server: ", err.Error())
	}
}
