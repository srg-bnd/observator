package agent

import (
	"log"

	config "github.com/srg-bnd/observator/config/agent"
	"github.com/srg-bnd/observator/internal/storage"
)

// Example demonstrates how to create and start a monitoring agent
func Example() {
	// Initialize the agent
	agent := NewAgent(&config.Container{
		ServerAddr:         "http://localhost:8080",
		Storage:            storage.NewMemStorage(),
		WorkerPoolReporter: 1,
	})

	// Start the agent with specified intervals
	err := agent.Start(60, 300)
	if err != nil {
		log.Fatal(err)
	}
}
