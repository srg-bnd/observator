// Package storage provides implementation for storing and managing metrics in various storage backends.
//
// The package supports both database and file-based storage solutions,
// allowing flexible configuration based on application requirements.
//
// Key features:
// - Metric storage and retrieval
// - Batch operations support
// - Multiple storage backends
// - Context-based operations
package storage

import (
	"context"
	"database/sql"
	"log"
)

// Repositories defines the interface for metric storage operations.
//
// This interface provides methods for:
// - Setting and getting gauge metrics
// - Setting and getting counter metrics
// - Batch operations for multiple metrics
// - Retrieving all stored metrics
//
// All methods are context-aware to support cancellation and timeouts.
type Repositories interface {
	SetGauge(context.Context, string, float64) error
	GetGauge(context.Context, string) (float64, error)
	SetCounter(context.Context, string, int64) error
	GetCounter(context.Context, string) (int64, error)
	SetBatchOfMetrics(context.Context, map[string]int64, map[string]float64) error
	AllCounterMetrics(context.Context) (map[string]int64, error)
	AllGaugeMetrics(context.Context) (map[string]float64, error)
}

type (
	// gauge represents a floating-point metric value
	gauge float64
	// counter represents an integer metric value
	counter int64
)

// Settings defines configuration parameters for the storage system.
//
// Fields:
// - DB: database connection
// - DatabaseDSN: connection string for database
// - FileStoragePath: path for file-based storage
// - StoreInterval: interval for periodic storage operations
// - Restore: flag to restore data from storage on startup
type Settings struct {
	DB              *sql.DB
	DatabaseDSN     string
	FileStoragePath string
	StoreInterval   int
	Restore         bool
}

// NewStorage creates a new storage instance based on provided settings.
//
// Parameters:
// - settings: configuration parameters for the storage system
//
// Returns:
// - Repositories: initialized storage implementation
//
// The function automatically selects the appropriate storage backend
// based on the provided configuration.
func NewStorage(settings *Settings) Repositories {
	var repStorage Repositories

	if settings.DatabaseDSN != "" {
		// DB Storage
		dbStorage := NewDBStorage(settings.DB)

		if err := dbStorage.ExecMigrations(); err != nil {
			log.Fatal(err)
		}

		repStorage = dbStorage
	} else {
		// File Storage
		fileStorage := NewFileStorage(settings.FileStoragePath, settings.StoreInterval, settings.Restore)
		if err := fileStorage.Load(); err != nil {
			log.Fatal(err)
		}
		fileStorage.Sync()
		repStorage = fileStorage
	}

	return repStorage
}
