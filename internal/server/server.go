package server

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/storage"
)

const (
	defaultHost = `:8080`
)

// Init server dependencies before startup
func Start(memStorage *storage.MemStorage) error {
	host, err := getHost()
	if err != nil {
		return err
	}

	http.ListenAndServe(host, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	return nil
}

// Selects the server host
func getHost() (string, error) {
	return defaultHost, nil
}
