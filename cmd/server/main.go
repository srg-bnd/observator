// Server that collects the metrics from agents
package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/router"
	"github.com/srg-bnd/observator/internal/storage"
)

// Application
type App struct {
	storage storage.Repositories
	server  *server.Server
	db      *sql.DB
}

// Returns a new application
func newApp() *App {
	db := newDB()
	storage := newStorage(db)

	return &App{
		storage: storage,
		server:  server.NewServer(router.NewRouter(storage, db)),
		db:      db,
	}
}

// Returns a new storage
func newStorage(db *sql.DB) storage.Repositories {
	var repStorage storage.Repositories

	if appConfigs.flagDatabaseDSN != "" {
		// DB Storage
		dbStorage := storage.NewDBStorage(db)

		if err := dbStorage.ExecMigrations(); err != nil {
			log.Fatal(err)
		}

		repStorage = dbStorage
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

// Returns a new connection to DB
func newDB() *sql.DB {
	db, err := sql.Open("pgx", appConfigs.flagDatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	parseFlags()

	// Init logger
	if err := logger.Initialize(appConfigs.flagLogLevel); err != nil {
		panic(err)
	}

	// Create application
	app := newApp()

	// Starts the application
	if err := app.server.Start(appConfigs.flagHostAddr); err != nil {
		log.Fatal(err)
	}

	// Close connection to DB
	defer app.db.Close()
}
