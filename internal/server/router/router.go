// Router
package router

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/middleware"
	"github.com/srg-bnd/observator/internal/storage"
)

// Returns a new router
func NewRouter(storage storage.Repositories, db *sql.DB) chi.Router {
	r := chi.NewRouter()

	showHandler := handlers.NewShowHandler(storage)
	updateHandler := handlers.NewUpdateHandler(storage)
	updatesHandler := handlers.NewUpdatesHandler(storage)

	r.Use(logger.RequestLogger, middleware.GzipMiddleware)
	r.Get("/ping", handlers.NewPingHandler(db).Handler)

	// Index
	r.Get("/", handlers.NewIndexHandler(storage).Handler)
	// Show
	r.Get("/value/{metricType}/{metricName}", showHandler.Handler)
	r.Post("/value", showHandler.JSONHandler)
	r.Post("/value/", showHandler.JSONHandler)
	// Update
	r.Post("/update/{metricType}/{metricName}/{metricValue}", updateHandler.Handler)
	r.Post("/update", updateHandler.JSONHandler)
	r.Post("/update/", updateHandler.JSONHandler)
	// Batch-update
	r.Post("/updates", updatesHandler.Handler)
	r.Post("/updates/", updatesHandler.Handler)

	return r
}
