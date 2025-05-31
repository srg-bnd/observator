// Handlers
package handlers

import (
	"net/http"
)

const (
	// Errors
	invalidDataError = "invalid data"
	invalidNameError = "invalid name"
	notFoundError    = "not found"
	serverError      = "server error"

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
	switch err.Error() {
	case serverError:
		w.WriteHeader(http.StatusInternalServerError)
	case invalidDataError:
		w.WriteHeader(http.StatusBadRequest)
	case notFoundError:
		w.WriteHeader(http.StatusNotFound)
	}
}
