// Flags & envs
package main

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

const (
	hostAddrUsage        = "address and port to run server"
	logLevellUsage       = "log level"
	storeIntervalUsage   = "store interval in seconds (zero for sync)"
	fileStoragePathUsage = "file storage path"
	restoreUsage         = "load data from storage"
	databaseDSNUsage     = "DB connection address"
	encryptionKeyUsage   = "encryption key"
)

const (
	hostAddrDefault        = ":8080"
	logLevellDefault       = "info"
	storeIntervalDefault   = 300
	fileStoragePathDefault = "./temp.storage.db"
	restoreDefault         = true
)

// Application configs
type AppConfigs struct {
	HostAddr string `env:"ADDRESS"`
	LogLevel string `env:"LOG_LEVEL"`
	// Storage
	StoreInterval   int    `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
	// Database
	DatabaseDSN string `env:"DATABASE_DSN"` // format: "host=%s user=%s password=%s dbname=%s sslmode=disable"
	// Encryption
	EncryptionKey string `env:"KEY"`
}

// Global app configs
var appConfigs = AppConfigs{}

// Parses flags from the console or envs
func parseFlags() {
	flag.StringVar(&appConfigs.HostAddr, "a", hostAddrDefault, hostAddrUsage)
	flag.StringVar(&appConfigs.LogLevel, "l", logLevellDefault, logLevellUsage)
	// Storage
	flag.IntVar(&appConfigs.StoreInterval, "i", storeIntervalDefault, storeIntervalUsage)
	flag.StringVar(&appConfigs.FileStoragePath, "f", fileStoragePathDefault, fileStoragePathUsage)
	flag.BoolVar(&appConfigs.Restore, "r", restoreDefault, restoreUsage)
	// Database
	flag.StringVar(&appConfigs.DatabaseDSN, "d", "", databaseDSNUsage)
	// Encryption
	flag.StringVar(&appConfigs.EncryptionKey, "k", "", encryptionKeyUsage)
	flag.Parse()

	err := env.Parse(&appConfigs)
	if err != nil {
		panic(err)
	}
}
