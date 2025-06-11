package serverHTTP

import (
	//"API/pkg/handler"
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer     *http.Server
	Host           string
	Port           string
	MaxHeaderBytes int
	Handler        http.Handler
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
}

func NewServer(host string, port string, maxHeaderBytes int, handler http.Handler, writeTimeout time.Duration, readTimeout time.Duration) *Server {
	return &Server{
		Host:           host,
		Port:           port,
		MaxHeaderBytes: maxHeaderBytes,
		Handler:        handler,
		WriteTimeout:   writeTimeout,
		ReadTimeout:    readTimeout,
	}
}

func (s *Server) Start() error {
	s.httpServer = &http.Server{
		Addr:           s.Host + ":" + s.Port,
		MaxHeaderBytes: s.MaxHeaderBytes,
		Handler:        s.Handler,
		WriteTimeout:   s.WriteTimeout,
		ReadTimeout:    s.ReadTimeout,
	}
	log.Printf("Server is running on %s:%s\n", s.Host, s.Port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Printf("Shutting down server on %s:%s\n", s.Host, s.Port)

	return s.httpServer.Shutdown(ctx)
}
