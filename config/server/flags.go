// Flags & envs
package config

import (
	"flag"

	"github.com/caarlos0/env/v11"
)

const (
	configUsage          = "config file"
	hostAddrUsage        = "base URL of target address"
	logLevellUsage       = "log level"
	storeIntervalUsage   = "store interval in seconds (zero for sync)"
	fileStoragePathUsage = "path to persistent file storage"
	restoreUsage         = "load data from storage"
	databaseDSNUsage     = "connection string to database"
	secretKeyUsage       = "sha256 key for hashing"
	cryptoKeyUsage       = "for asymmetric asymmetric"
)

const (
	hostAddrDefault        = ":8080"
	logLevellDefault       = "info"
	storeIntervalDefault   = 300
	fileStoragePathDefault = "./temp.storage.db"
	restoreDefault         = true
)

// Application configs
type flags struct {
	ConfigFile string `env:"CONFIG"`

	HostAddr string `env:"ADDRESS"`
	LogLevel string `env:"LOG_LEVEL"`
	// Storage
	StoreInterval   int    `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
	// Database
	DatabaseDSN string `env:"DATABASE_DSN"` // format: "host=%s user=%s password=%s dbname=%s sslmode=disable"
	// Encryption
	SecretKey string `env:"KEY"`
	CryptoKey string `env:"CRYPTO_KEY"`
}

// Global app configs
var Flags = flags{}

// Parses flags from the console or envs
func (s *flags) ParseFlags() {
	flag.StringVar(&Flags.ConfigFile, "c", "", configUsage)
	// TODO: parse ConfigFile to flags struct

	flag.StringVar(&Flags.HostAddr, "a", hostAddrDefault, hostAddrUsage)
	flag.StringVar(&Flags.LogLevel, "l", logLevellDefault, logLevellUsage)
	// Storage
	flag.IntVar(&Flags.StoreInterval, "i", storeIntervalDefault, storeIntervalUsage)
	flag.StringVar(&Flags.FileStoragePath, "f", fileStoragePathDefault, fileStoragePathUsage)
	flag.BoolVar(&Flags.Restore, "r", restoreDefault, restoreUsage)
	// Database
	flag.StringVar(&Flags.DatabaseDSN, "d", "", databaseDSNUsage)
	// Encryption
	flag.StringVar(&Flags.SecretKey, "k", "", secretKeyUsage)
	flag.StringVar(&Flags.CryptoKey, "crypto-key", "", cryptoKeyUsage)
	flag.Parse()

	err := env.Parse(&Flags)
	if err != nil {
		panic(err)
	}
}
