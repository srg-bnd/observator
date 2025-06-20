// Handlers
package handlers

import (
	"errors"
	"net/http"
)

var (
	// Errors
	errInvalidData = errors.New("invalid data")
	errInvalidName = errors.New("invalid name")
	errNotFound    = errors.New("not found")
	errServer      = errors.New("server error")
)

const (
	// Formats
	JSONFormat = "json"
	HTMLFormat = "text/html"
	TextFormat = "text"
)

// Set content type
func setContentType(w http.ResponseWriter, format string) {
	switch format {
	case JSONFormat:
		w.Header().Set("Content-Type", "application/json")
	case HTMLFormat:
		w.Header().Set("Content-Type", "text/html")
	case TextFormat:
		w.Header().Set("Content-Type", "text/plain")
	}
}

// Handle errors
func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, errNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, errInvalidData):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
