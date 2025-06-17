// Flags & envs
package main

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

const (
	pollIntervalUsage   = "pollInterval – frequency (seconds) of metric polling"
	reportIntervalUsage = "reportInterval – frequency (seconds) of sending values to the server"
	serverAddrUsage     = "address and port to run server"
	secretKeyUsage      = "encryption key"
)

const (
	pollIntervalDefault   = 2
	reportIntervalDefault = 10
	serverAddrDefault     = "localhost:8080"
)

// Application configs
type AppConfigs struct {
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	ServerAddr     string `env:"ADDRESS"`
	SecretKey      string `env:"KEY"`
}

// Global app configs
var appConfigs = AppConfigs{}

// Parses flags & envs
func parseFlags() {
	flag.IntVar(&appConfigs.PollInterval, "p", pollIntervalDefault, pollIntervalUsage)
	flag.IntVar(&appConfigs.ReportInterval, "r", reportIntervalDefault, reportIntervalUsage)
	flag.StringVar(&appConfigs.ServerAddr, "a", serverAddrDefault, serverAddrUsage)
	flag.StringVar(&appConfigs.SecretKey, "k", "", secretKeyUsage)
	flag.Parse()

	err := env.Parse(&appConfigs)
	if err != nil {
		panic(err)
	}
}
