// Server
package server

import (
	"net/http"
)

// Base server
type Server struct {
	router http.Handler
}

// Creates a new server
func NewServer(router http.Handler) *Server {
	return &Server{
		router: router,
	}
}

// Starts the server
func (server *Server) Start(addr string) error {
	err := http.ListenAndServe(addr, server.router)
	if err != nil {
		return err
	}

	return nil
}
