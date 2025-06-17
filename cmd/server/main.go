// Server that collects the metrics from agents
package main

import (
	"log"

	config "github.com/srg-bnd/observator/config/server"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/db"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/router"
	"github.com/srg-bnd/observator/internal/storage"
)

func main() {
	parseFlags()

	// Init global logger
	if err := logger.Initialize(appConfigs.flagLogLevel); err != nil {
		panic(err)
	}

	// Init DI Container
	db := db.NewDB(appConfigs.flagDatabaseDSN)
	container := &config.Container{
		DB: db,
		Storage: storage.NewStorage(
			&storage.Settings{
				DB:                  db,
				FlagDatabaseDSN:     appConfigs.flagDatabaseDSN,
				FlagFileStoragePath: appConfigs.flagFileStoragePath,
				FlagStoreInterval:   appConfigs.flagStoreInterval,
				FlagRestore:         appConfigs.flagRestore,
			},
		),
	}

	// Starts the server
	if err := server.NewServer(router.NewRouter(container)).Start(appConfigs.flagHostAddr); err != nil {
		log.Fatal(err)
	}

	// Close connection to DB
	defer db.Close()
}
