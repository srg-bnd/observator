// Flags
package main

import (
	"flag"
	"os"
	"strconv"
)

type AppConfigs struct {
	flagHostAddr        string
	flagLogLevel        string
	flagStoreInterval   int
	flagFileStoragePath string
	flagRestore         bool
}

var appConfigs = AppConfigs{}

func parseFlags() {
	flag.StringVar(&appConfigs.flagHostAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&appConfigs.flagLogLevel, "l", "info", "log level")
	// Files
	flag.IntVar(&appConfigs.flagStoreInterval, "i", 300, "store interval in seconds (zero for sync)")
	flag.StringVar(&appConfigs.flagFileStoragePath, "f", "~/temp.storage", "file storage path")
	flag.BoolVar(&appConfigs.flagRestore, "r", true, "load data from storage")
	flag.Parse()

	if envHostAddr := os.Getenv("ADDRESS"); envHostAddr != "" {
		appConfigs.flagHostAddr = envHostAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		appConfigs.flagLogLevel = envLogLevel
	}
	// Files
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
}
