// Server for metrics collection and alerting service
package main

import (
	"log"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/storage"
)

type App struct {
	storage *storage.MemStorage
	server  *server.Server
}

func NewApp() *App {
	storage := storage.NewMemStorage(appConfigs.flagFileStoragePath, appConfigs.flagStoreInterval, appConfigs.flagRestore)

	return &App{
		storage: storage,
		server:  server.NewServer(storage),
	}
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
}
