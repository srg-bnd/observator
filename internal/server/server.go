// Server
package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/storage"
)

type Server struct {
	handler Handler
}

type Handler interface {
	GetRouter() chi.Router
}

// Creates a new server
func NewServer(storage storage.Repositories, db *sql.DB) *Server {
	return &Server{
		handler: handlers.NewHandler(storage, db),
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
