// Server
package server

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/storage"
)

type Server struct {
	handler *handlers.Handler
}

// Creates a new server
func NewServer(storage storage.Repositories) *Server {
	return &Server{
		handler: handlers.NewHandler(storage),
	}
}

// Starts the server
func (server *Server) Start(addr string) error {
	err := http.ListenAndServe(addr, server.handler.GetRouter())
	if err != nil {
		return err
	}

	return nil
}
