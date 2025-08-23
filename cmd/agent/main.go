// Agent that collects the metrics and sends them to the server
package main

import (
	"fmt"
	"log"

	config "github.com/srg-bnd/observator/config/agent"
	"github.com/srg-bnd/observator/internal/agent"
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
	// Init DI container
	container := &config.Container{
		Storage:            storage.NewMemStorage(),
		ServerAddr:         config.Flags.ServerAddr,
		WorkerPoolReporter: config.Flags.RateLimit,
	}

	if config.Flags.SecretKey != "" {
		container.ChecksumService = services.NewChecksum(config.Flags.SecretKey)
	}

	if err := agent.NewAgent(container).Start(config.Flags.PollInterval, config.Flags.ReportInterval); err != nil {
		log.Fatal(err)
	}
}
