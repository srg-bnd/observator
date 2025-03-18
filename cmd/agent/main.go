// Agent for metrics collection and alerting service
package main

import (
	"log"

	"github.com/srg-bnd/observator/internal/agent"
	"github.com/srg-bnd/observator/internal/storage"
)

type App struct {
	agent *agent.Agent
}

func NewApp() *App {
	// HACK
	storage := storage.NewMemStorage("", 0, false)
	return &App{
		agent: agent.NewAgent(storage, appConfigs.serverAddr),
	}
}

func main() {
	parseFlags()

	if err := NewApp().agent.Start(appConfigs.pollInterval, appConfigs.reportInterval); err != nil {
		log.Fatal(err)
	}
}
