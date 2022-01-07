package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	server *http.Server
}

func Get() *Server {
	return &Server{
		server: &http.Server{},
	}
}

func (s *Server) WithAddress(address string) *Server {
	s.server.Addr = address
	return s
}

func (s *Server) WithErrorLogger(logger *log.Logger) *Server {
	s.server.ErrorLog = logger
	return s
}

func (s *Server) WithRouter(router *httprouter.Router) *Server {
	s.server.Handler = router
	return s
}

func (s *Server) Start() error {
	if len(s.server.Addr) == 0 {
		return errors.New("server missing address")
	}

	if s.server.Handler == nil {
		return errors.New("server missing handler")
	}

	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}
