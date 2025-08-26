// Server that collects the metrics from agents
package main

import (
	"fmt"
	"log"

	config "github.com/srg-bnd/observator/config/server"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/server/db"
	"github.com/srg-bnd/observator/internal/server/logger"
	"github.com/srg-bnd/observator/internal/server/router"
	"github.com/srg-bnd/observator/internal/shared/services"
	"github.com/srg-bnd/observator/internal/storage"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func buildInfo() {
	strOrNA := func(str string) string {
		if str == "" {
			return "N/A"
		}

		return str
	}

	fmt.Printf(
		"Build version: %s\nBuild date: %s\nBuild commit: %s\n",
		strOrNA(buildVersion),
		strOrNA(buildDate),
		strOrNA(buildCommit),
	)
}

func init() {
	buildInfo()
	config.Flags.ParseFlags()
}

func main() {
	// Init global logger
	if err := logger.Initialize(config.Flags.LogLevel); err != nil {
		panic(err)
	}

	var checksumService *services.Checksum
	if config.Flags.SecretKey != "" {
		checksumService = services.NewChecksum(config.Flags.SecretKey)
	}

	var privateKey *services.PrivateKey
	if config.Flags.CryptoKey != "" {
		privateKey = services.NewPrivateKey(config.Flags.CryptoKey)
	}

	// Init DI Container
	db := db.NewDB(config.Flags.DatabaseDSN)
	container := &config.Container{
		DB:              db,
		ChecksumService: checksumService,
		PrivateKey:      privateKey,
		Storage: storage.NewStorage(
			&storage.Settings{
				DB:              db,
				DatabaseDSN:     config.Flags.DatabaseDSN,
				FileStoragePath: config.Flags.FileStoragePath,
				StoreInterval:   config.Flags.StoreInterval,
				Restore:         config.Flags.Restore,
			},
		),
	}

	// Starts the server
	if err := server.NewServer(router.NewRouter(container)).Start(config.Flags.HostAddr); err != nil {
		log.Fatal(err)
	}

	// Close connection to DB
	defer db.Close()
}
