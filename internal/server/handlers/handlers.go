// Handlers & Router for server
package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/services"
	"github.com/srg-bnd/observator/internal/storage"
)

const (
	serverError      = "serverError"
	invalidDataError = "invalidDataError"
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

	r.Get("/", logger.RequestLogger(h.IndexHandler))
	r.Get("/value/{metricType}/{metricName}", logger.RequestLogger(h.ShowHandler))
	r.Post("/value", logger.RequestLogger(h.ShowAsJSONHandler))
	r.Post("/value/", logger.RequestLogger(h.ShowAsJSONHandler))
	r.Post("/update/{metricType}/{metricName}/{metricValue}", logger.RequestLogger(h.UpdateHandler))
	r.Post("/update", logger.RequestLogger(h.UpdateAsJSONHandler))

	return r
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
