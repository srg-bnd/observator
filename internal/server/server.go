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

func NewServer(storage storage.Repositories) *Server {
	return &Server{
		handler: handlers.NewHandler(storage),
	}
}

// Init server dependencies before startup
func (server *Server) Start() error {
	host, err := getHost()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle(
		`/`,
		middleware.Chain(
			http.HandlerFunc(server.handler.IndexHandler),
		))

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

	http.ListenAndServe(host, mux)

	return nil
}

// Selects the server host
func getHost() (string, error) {
	return defaultHost, nil
}
