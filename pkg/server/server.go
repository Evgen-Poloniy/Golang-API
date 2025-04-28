package serverHTTP

import (
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
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
}

func NewServer(host string, port string, maxHeaderBytes int, writeTimeout time.Duration, readTimeout time.Duration) *Server {
	return &Server{
		Host:           host,
		Port:           port,
		MaxHeaderBytes: maxHeaderBytes,
		WriteTimeout:   writeTimeout,
		ReadTimeout:    readTimeout,
	}
}

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	var resp Response = Response{
		Status:  http.StatusOK,
		Message: "Connection successful",
	}

	jsonResponce(resp, w)

	var address string = r.RemoteAddr
	log.Printf("Ping: Connection successful, Address: %s", address)
}

func (s *Server) Start() error {
	s.httpServer = &http.Server{
		Addr:           s.Host + ":" + s.Port,
		MaxHeaderBytes: s.MaxHeaderBytes,
		WriteTimeout:   s.WriteTimeout,
		ReadTimeout:    s.ReadTimeout,
	}
	log.Printf("Server is running on %s:%s", s.Host, s.Port)

	http.HandleFunc("/ping", s.Ping)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Printf("Shutting down server on %s:%s...", s.Host, s.Port)

	return s.httpServer.Shutdown(ctx)
}
