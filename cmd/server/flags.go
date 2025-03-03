// Flags
package main

import (
	"flag"
	"os"
)

type AppConfigs struct {
	flagHostAddr string
	flagLogLevel string
}

var appConfigs = AppConfigs{}

func parseFlags() {
	flag.StringVar(&appConfigs.flagHostAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&appConfigs.flagLogLevel, "l", "info", "log level")
	flag.Parse()

	if envHostAddr := os.Getenv("ADDRESS"); envHostAddr != "" {
		appConfigs.flagHostAddr = envHostAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		appConfigs.flagLogLevel = envLogLevel
	}
}
