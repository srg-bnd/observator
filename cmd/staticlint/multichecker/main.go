package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

// Config is the name of the configuration file
const Config = `config.json`

// ConfigData describes the structure of the configuration file
type ConfigData struct {
	Staticheck []string
}

func getConfig() *ConfigData {
	appfile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(filepath.Join(filepath.Dir(appfile), Config))
	if err != nil {
		panic(err)
	}
	var cfg ConfigData
	if err = json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}

func main() {
	cfg := getConfig()
	var analyzers []*analysis.Analyzer

	// Adds standard analyzers
	analyzers = append(analyzers, standardAnalyzers()...)
	// Adds staticcheck analyzers
	analyzers = append(analyzers, staticcheckAnalyzers(cfg.Staticheck)...)
	// Adds vendor analyzers
	analyzers = append(analyzers, vendorAnalyzers()...)
	// Adds custom analyzers
	analyzers = append(analyzers, customAnalyzers()...)

	// Runs multichecker
	multichecker.Main(analyzers...)
}
