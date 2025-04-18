// Metrics collection and alerting service
package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/storage"
)

// App
type App struct {
	storage storage.Repositories
	server  *server.Server
	db      *sql.DB
}

// Returns App
func newApp() *App {
	db := newDB()
	storage := newStorage(db)

	return &App{
		storage: storage,
		server:  server.NewServer(storage, db),
		db:      db,
	}
}

// Returns Storage
func newStorage(db *sql.DB) storage.Repositories {
	var repStorage storage.Repositories

	// DB Storage
	if appConfigs.flagDatabaseDSN != "" {
		dbStorage := storage.NewDBStorage(db)
		repStorage = dbStorage
		// Migrations
		migrationPath := "./db/migrations"
		entries, err := os.ReadDir(migrationPath)
		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			query, err := os.ReadFile(migrationPath + "/" + entry.Name())
			if err != nil {
				panic(err)
			}

			_, err = db.ExecContext(context.Background(), string(query))
			if err != nil {
				panic(err)
			}
		}
	} else {
		// File Storage
		fileStorage := storage.NewFileStorage(appConfigs.flagFileStoragePath, appConfigs.flagStoreInterval, appConfigs.flagRestore)
		if err := fileStorage.Load(); err != nil {
			log.Fatal(err)
		}
		fileStorage.Sync()
		repStorage = fileStorage
	}

	return repStorage
}

// Returns DB
func newDB() *sql.DB {
	db, err := sql.Open("pgx", appConfigs.flagDatabaseDSN)
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	// Parse run-flags
	parseFlags()

	// Init logger
	if err := logger.Initialize(appConfigs.flagLogLevel); err != nil {
		panic(err)
	}

	// Create App
	app := newApp()

	// Start App
	if err := app.server.Start(appConfigs.flagHostAddr); err != nil {
		log.Fatal(err)
	}

	// Close DB
	defer app.db.Close()
}
