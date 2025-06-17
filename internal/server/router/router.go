// Router
package router

import (
	config "github.com/srg-bnd/observator/config/server"

	"github.com/go-chi/chi"
	"github.com/srg-bnd/observator/internal/server/handlers"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/middleware"
)

// Returns a new router
func NewRouter(container *config.Container) chi.Router {
	r := chi.NewRouter()

	showHandler := handlers.NewShowHandler(container.Storage)
	updateHandler := handlers.NewUpdateHandler(container.Storage)
	batchUpdateHandler := handlers.NewBatchUpdateHandler(container.Storage)

	r.Use(logger.RequestLogger, middleware.GzipMiddleware)

	if container.EncryptionKey != "" {
		r.Use(middleware.NewChecksum(container.EncryptionKey).WithVerify)
	}

	r.Get("/ping", handlers.NewPingHandler(container.DB).Handler)

	// Index
	r.Get("/", handlers.NewIndexHandler(container.Storage).Handler)
	// Show
	r.Get("/value/{metricType}/{metricName}", showHandler.Handler)
	r.Post("/value", showHandler.JSONHandler)
	r.Post("/value/", showHandler.JSONHandler)
	// Update
	r.Post("/update/{metricType}/{metricName}/{metricValue}", updateHandler.Handler)
	r.Post("/update", updateHandler.JSONHandler)
	r.Post("/update/", updateHandler.JSONHandler)
	// Batch-update
	r.Post("/updates", batchUpdateHandler.Handler)
	r.Post("/updates/", batchUpdateHandler.Handler)

	return r
}
