// Flags & envs
package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

const (
	configUsage         = "config file"
	pollIntervalUsage   = "pollInterval – frequency (seconds) of metric polling"
	rateLimitUsage      = "rate limit"
	reportIntervalUsage = "reportInterval – frequency (seconds) of sending values to the server"
	secretKeyUsage      = "encryption key"
	serverAddrUsage     = "address and port to run server"
	cryptoKeyUsage      = "for asymmetric asymmetric"
)

const (
	pollIntervalDefault   = 2
	rateLimitDefault      = 1
	reportIntervalDefault = 10
	serverAddrDefault     = "localhost:8080"
)

// Application configs
type flags struct {
	ConfigFile string `env:"CONFIG"`

	PollInterval   int    `env:"POLL_INTERVAL"`
	RateLimit      int    `env:"RATE_LIMIT"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	SecretKey      string `env:"KEY"`
	ServerAddr     string `env:"ADDRESS"`
	CryptoKey      string `env:"CRYPTO_KEY"`
}

// Global app configs
var Flags = flags{}

// Parses flags & envs
func (s *flags) ParseFlags() {
	flag.StringVar(&Flags.ConfigFile, "c", "", configUsage)
	// TODO: parse ConfigFile to flags struct

	flag.IntVar(&Flags.PollInterval, "p", pollIntervalDefault, pollIntervalUsage)
	flag.IntVar(&Flags.RateLimit, "l", rateLimitDefault, rateLimitUsage)
	flag.IntVar(&Flags.ReportInterval, "r", reportIntervalDefault, reportIntervalUsage)
	flag.StringVar(&Flags.SecretKey, "k", "", secretKeyUsage)
	flag.StringVar(&Flags.ServerAddr, "a", serverAddrDefault, serverAddrUsage)
	flag.StringVar(&Flags.CryptoKey, "crypto-key", "", cryptoKeyUsage)
	flag.Parse()

	err := env.Parse(&Flags)
	if err != nil {
		panic(err)
	}
}
