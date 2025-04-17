// Metrics collection and alerting service
package main

import (
	"database/sql"
	"log"

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
	storage := newStorage()
	db := newDB()

	return &App{
		storage: storage,
		server:  server.NewServer(storage, db),
		db:      db,
	}
}

// Returns Storage
func newStorage() storage.Repositories {
	// TODO: DBStorage (-d present)

	// FileStorage (-f present)
	storage := storage.NewFileStorage(appConfigs.flagFileStoragePath, appConfigs.flagStoreInterval, appConfigs.flagRestore)
	if err := storage.Load(); err != nil {
		log.Fatal(err)
	}
	storage.Sync()

	// TODO: MemStorage (else)

	return storage
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
