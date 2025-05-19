// Handlers & Router for server
package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/middleware"
	"github.com/srg-bnd/observator/internal/storage"
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

type Handler struct {
	storage storage.Repositories
	db      *sql.DB
}

// Returns new Handler
func NewHandler(storage storage.Repositories, db *sql.DB) *Handler {
	return &Handler{
		storage: storage,
		db:      db,
	}
}

// Returns Router for HTTP Server
func (h *Handler) GetRouter() chi.Router {
	r := chi.NewRouter()

	showHandler := NewShowHandler(h.storage)
	updateHandler := NewUpdateHandler(h.storage)
	updatesHandler := NewUpdatesHandler(h.storage)

	r.Use(logger.RequestLogger, middleware.GzipMiddleware)
	r.Get("/ping", h.PingHandler)

	// Index
	r.Get("/", NewIndexHandler(h.storage).Handler)
	// Show
	r.Get("/value/{metricType}/{metricName}", showHandler.Handler)
	r.Post("/value", showHandler.JSONHandler)
	r.Post("/value/", showHandler.JSONHandler)
	// Single-update
	r.Post("/update/{metricType}/{metricName}/{metricValue}", updateHandler.Handler)
	r.Post("/update", updateHandler.JSONHandler)
	r.Post("/update/", updateHandler.JSONHandler)
	// Batch-update
	r.Post("/updates", updatesHandler.Handler)
	r.Post("/updates/", updatesHandler.Handler)

	return r
}

// GET /ping
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := h.db.PingContext(ctx); err != nil {
		log.Println("error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
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
	case notFoundError:
		w.WriteHeader(http.StatusNotFound)
	}
}
