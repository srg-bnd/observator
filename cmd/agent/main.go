// Agent for metrics collection and alerting service
package main

import (
	"log"

	"github.com/srg-bnd/observator/internal/agent"
)

func main() {
	if err := agent.Start(); err != nil {
		log.Fatal(err)
	}
}
