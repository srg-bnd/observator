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
	storage := storage.NewMemStorage()
	return &App{
		agent: agent.NewAgent(storage, AgentFlags.serverAddr),
	}
}

func main() {
	parseFlags()

	if err := NewApp().agent.Start(AgentFlags.pollInterval, AgentFlags.reportInterval); err != nil {
		log.Fatal(err)
	}
}
