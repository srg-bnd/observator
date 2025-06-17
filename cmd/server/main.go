// Server that collects the metrics from agents
package main

import (
	"log"

	config "github.com/srg-bnd/observator/config/server"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/db"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/router"
	"github.com/srg-bnd/observator/internal/shared/services"
	"github.com/srg-bnd/observator/internal/storage"
)

func main() {
	parseFlags()

	// Init global logger
	if err := logger.Initialize(appConfigs.LogLevel); err != nil {
		panic(err)
	}

	var checksumService *services.Checksum
	if appConfigs.SecretKey != "" {
		checksumService = services.NewChecksum(appConfigs.SecretKey)
	}

	// Init DI Container
	db := db.NewDB(appConfigs.DatabaseDSN)
	container := &config.Container{
		DB:              db,
		ChecksumService: checksumService,
		Storage: storage.NewStorage(
			&storage.Settings{
				DB:              db,
				DatabaseDSN:     appConfigs.DatabaseDSN,
				FileStoragePath: appConfigs.FileStoragePath,
				StoreInterval:   appConfigs.StoreInterval,
				Restore:         appConfigs.Restore,
			},
		),
	}

	// Starts the server
	if err := server.NewServer(router.NewRouter(container)).Start(appConfigs.HostAddr); err != nil {
		log.Fatal(err)
	}

	// Close connection to DB
	defer db.Close()
}
