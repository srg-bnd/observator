package server

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/handlers"
	"github.com/srg-bnd/observator/internal/middleware"
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
		`/update/{metricType}/{metricName}/{metricValue}`,
		middleware.Chain(
			http.HandlerFunc(server.handler.UpdateMetricHandler),
			middleware.CheckMethodPost,
		))

	http.ListenAndServe(host, mux)

	return nil
}

// Selects the server host
func getHost() (string, error) {
	return defaultHost, nil
}
