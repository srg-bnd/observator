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
	server.MemStorage = &memStorage
}
func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
