// Server
package server

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultHost = `:8080`
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
func (server *Server) Start() error {
	host, err := getHost()
	if err != nil {
		return err
	}

	err = http.ListenAndServe(host, server.handler.GetRouter())
	if err != nil {
		return err
	}

	return nil
}

// Selects the server host
func getHost() (string, error) {
	return defaultHost, nil
}
