// Server for metrics collection and alerting service
package main

import (
	"log"

	"github.com/srg-bnd/observator/internal/server"
	"github.com/srg-bnd/observator/internal/storage"
)

var memStorage storage.MemStorage

func init() {
	memStorage = storage.NewMemStorage()
}

func main() {
	if err := server.Start(&memStorage); err != nil {
		log.Fatal(err)
	}
}
