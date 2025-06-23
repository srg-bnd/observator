// Storage for metrics
package storage

import (
	"context"
	"database/sql"
	"log"
)

type Repositories interface {
	SetGauge(context.Context, string, float64) error
	GetGauge(context.Context, string) (float64, error)
	SetCounter(context.Context, string, int64) error
	GetCounter(context.Context, string) (int64, error)
	SetBatchOfMetrics(context.Context, map[string]int64, map[string]float64) error
}

type (
	gauge   float64
	counter int64
)

type Settings struct {
	DB              *sql.DB
	DatabaseDSN     string
	FileStoragePath string
	StoreInterval   int
	Restore         bool
}

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
