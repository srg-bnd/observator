// Handlers
package handlers

import (
	"errors"
	"net/http"
)

var (
	// Errors
	invalidDataError = errors.New("invalid data")
	invalidNameError = errors.New("invalid name")
	notFoundError    = errors.New("not found")
	serverError      = errors.New("server error")
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
	case errors.Is(err, notFoundError):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, invalidDataError):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
