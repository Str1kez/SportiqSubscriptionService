package server

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Run(host, port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           host + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return server.httpServer.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	log.Infoln("shutting down...")
	return server.httpServer.Shutdown(ctx)
}
