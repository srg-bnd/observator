// Handlers & Router for server
package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/middleware"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	// Errors
	serverError      = "serverError"
	invalidDataError = "invalidDataError"

	// Formats
	JSONFormat = "json"
	HTMLFormat = "text/html"
	TextFormat = "text"
)

type Handler struct {
	service *services.Service
	storage storage.Repositories
}

// Returns new Handler
func NewHandler(storage storage.Repositories) *Handler {
	return &Handler{
		service: services.NewService(storage),
		storage: storage,
	}
}

// Returns Router for HTTP Server
func (h *Handler) GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(logger.RequestLogger, middleware.GzipMiddleware)
	r.Get("/", h.IndexHandler)
	r.Get("/value/{metricType}/{metricName}", h.ShowHandler)
	r.Post("/value", h.ShowAsJSONHandler)
	r.Post("/value/", h.ShowAsJSONHandler)
	r.Post("/update/{metricType}/{metricName}/{metricValue}", h.UpdateHandler)
	r.Post("/update", h.UpdateAsJSONHandler)
	r.Post("/update/", h.UpdateAsJSONHandler)

	return r
}

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
	}

}
