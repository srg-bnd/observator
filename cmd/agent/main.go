// Agent that collects the metrics and sends them to the server
package main

import (
	"log"

	config "github.com/srg-bnd/observator/config/agent"
	"github.com/srg-bnd/observator/internal/agent"
	"github.com/srg-bnd/observator/internal/storage"
)

func main() {
	parseFlags()

	// Init DI container
	container := &config.Container{
		Storage:    storage.NewMemStorage(),
		ServerAddr: appConfigs.ServerAddr,
	}

	if err := agent.NewAgent(container).Start(appConfigs.PollInterval, appConfigs.ReportInterval); err != nil {
		log.Fatal(err)
	}
}
