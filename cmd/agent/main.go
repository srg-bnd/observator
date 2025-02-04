// Agent for metrics collection and alerting service
package main

import (
	"log"

	"github.com/srg-bnd/observator/internal/agent"
	"github.com/srg-bnd/observator/internal/storage"
)

var memStorage storage.MemStorage

func init() {
	memStorage = storage.NewMemStorage()
	// HACK to acess memStorage
	agent.MemStorage = &memStorage
}

func main() {
	if err := agent.Start(); err != nil {
		log.Fatal(err)
	}
}
