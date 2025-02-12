package server

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/handlers"
	"github.com/srg-bnd/observator/internal/middleware"
)

const (
	defaultHost = `:8080`
)

var (
	mux *http.ServeMux
)

func init() {
	// Routes
	mux = http.NewServeMux()
	mux.Handle(
		`/update/{metricType}/{metricName}/{metricValue}`,
		middleware.Chain(
			http.HandlerFunc(handlers.UpdateMetricHandler),
			middleware.CheckMethodPost,
		))
}

// Init server dependencies before startup
func Start() error {
	host, err := getHost()
	if err != nil {
		return err
	}

	http.ListenAndServe(host, mux)

	return nil
}

// Selects the server host
func getHost() (string, error) {
	return defaultHost, nil
}
