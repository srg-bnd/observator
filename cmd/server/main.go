// Server for metrics collection and alerting service
package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/storage"
)

type App struct {
	storage *storage.MemStorage
	server  *server.Server
	db      *sql.DB
}

func NewApp() *App {
	storage := storage.NewMemStorage(appConfigs.flagFileStoragePath, appConfigs.flagStoreInterval, appConfigs.flagRestore)
	db := newDB()

	return &App{
		storage: storage,
		server:  server.NewServer(storage, db),
		db:      db,
	}
}

func newDB() *sql.DB {
	db, err := sql.Open("pgx", appConfigs.flagDatabaseDSN)
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	parseFlags()

	if err := logger.Initialize(appConfigs.flagLogLevel); err != nil {
		panic(err)
	}

	app := NewApp()
	if err := app.storage.Load(); err != nil {
		log.Fatal(err)
	}
	app.storage.Sync()

	if err := app.server.Start(appConfigs.flagHostAddr); err != nil {
		log.Fatal(err)
	}

	defer app.db.Close()
}
