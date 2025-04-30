// Agent that collects the metrics and sends them to the server
package main

import (
	"log"

	"github.com/srg-bnd/observator/internal/agent"
	"github.com/srg-bnd/observator/internal/storage"
)

// Application
type App struct {
	agent *agent.Agent
}

// Returns a new application
func NewApp() *App {
	return &App{
		agent: agent.NewAgent(storage.NewMemStorage(), appConfigs.serverAddr),
	}
}

func main() {
	parseFlags()

	if err := NewApp().agent.Start(appConfigs.pollInterval, appConfigs.reportInterval); err != nil {
		log.Fatal(err)
	}
}
