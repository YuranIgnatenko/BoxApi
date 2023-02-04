package server

import (
	"BoxApi/config"
)

type Server struct {
	Host string
	Port string
}

func NewServer(config config.Config) *Server {
	return &Server{
		Host: config["Host"],
		Port: config["Port"],
	}
}

func (s *Server) Addr() string {
	return s.Host + ":" + s.Port
}
