package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server   *http.Server
	shutdown time.Duration
}

func New(handler http.Handler, p string) *Server {

	server1 := &http.Server{
		Addr:         p,
		Handler:      handler,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	s := &Server{
		server:   server1,
		shutdown: 5 * time.Second,
	}

	return s
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdown)
	defer cancel()

	return s.server.Shutdown(ctx)
}
