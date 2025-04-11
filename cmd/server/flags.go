// Flags
package main

import (
	"flag"
	"os"
	"strconv"
)

type AppConfigs struct {
	flagHostAddr string
	flagLogLevel string
	// Storage
	flagStoreInterval   int
	flagFileStoragePath string
	flagRestore         bool
	// Database
	flagDatabaseDSN string // "host=%s user=%s password=%s dbname=%s sslmode=disable"
}

var appConfigs = AppConfigs{}

func parseFlags() {
	flag.StringVar(&appConfigs.flagHostAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&appConfigs.flagLogLevel, "l", "info", "log level")
	// Storage
	flag.IntVar(&appConfigs.flagStoreInterval, "i", 300, "store interval in seconds (zero for sync)")
	flag.StringVar(&appConfigs.flagFileStoragePath, "f", "./temp.storage.db", "file storage path")
	flag.BoolVar(&appConfigs.flagRestore, "r", true, "load data from storage")
	// Database
	flag.StringVar(&appConfigs.flagDatabaseDSN, "d", "", "DB connection address")
	flag.Parse()

	if envHostAddr := os.Getenv("ADDRESS"); envHostAddr != "" {
		appConfigs.flagHostAddr = envHostAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		appConfigs.flagLogLevel = envLogLevel
	}
	// Storage
	if envStoreInterval := os.Getenv("STORE_INTERVAL"); envStoreInterval != "" {
		value, _ := strconv.ParseInt(envStoreInterval, 10, 0)
		appConfigs.flagStoreInterval = int(value)
	}
	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		appConfigs.flagFileStoragePath = envFileStoragePath
	}
	if envRestore := os.Getenv("RESTORE"); envRestore != "" {
		value, _ := strconv.ParseBool(envRestore)
		appConfigs.flagRestore = value
	}
	// Database
	if envDatabaseDSN := os.Getenv("DATABASE_DSN"); envDatabaseDSN != "" {
		appConfigs.flagDatabaseDSN = envDatabaseDSN
	}
}
