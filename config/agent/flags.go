// Flags & envs
package config

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
type flags struct {
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	ServerAddr     string `env:"ADDRESS"`
	SecretKey      string `env:"KEY"`
}

// Global app configs
var Flags = flags{}

// Parses flags & envs
func (s *flags) ParseFlags() {
	flag.IntVar(&Flags.PollInterval, "p", pollIntervalDefault, pollIntervalUsage)
	flag.IntVar(&Flags.ReportInterval, "r", reportIntervalDefault, reportIntervalUsage)
	flag.StringVar(&Flags.ServerAddr, "a", serverAddrDefault, serverAddrUsage)
	flag.StringVar(&Flags.SecretKey, "k", "", secretKeyUsage)
	flag.Parse()

	err := env.Parse(&Flags)
	if err != nil {
		panic(err)
	}
}
