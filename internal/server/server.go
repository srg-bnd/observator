package server

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/server/middleware"
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

	err = http.ListenAndServe(host, server.GetMux())
	if err != nil {
		return err
	}

	return nil
}

// Init router
func (server *Server) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(
		`/update/{metricType}/{metricName}/{metricValue}`,
		middleware.Chain(
			http.HandlerFunc(server.handler.UpdateMetricHandler),
			middleware.CheckMethodPost,
		))

	mux.Handle(
		`/value/{metricType}/{metricName}`,
		middleware.Chain(
			http.HandlerFunc(server.handler.ShowMetricHandler),
		))

	mux.Handle(
		`/`,
		middleware.Chain(
			http.HandlerFunc(server.handler.IndexHandler),
		))

	return mux
}

// Selects the server host
func getHost() (string, error) {
	return defaultHost, nil
}
