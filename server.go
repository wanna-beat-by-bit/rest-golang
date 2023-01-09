package rsapi

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, hanlder http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        hanlder,
		MaxHeaderBytes: 1 << 20, //1 mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
