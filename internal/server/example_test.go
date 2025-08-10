package server

import (
	"log"

	"github.com/go-resty/resty/v2"
	config "github.com/srg-bnd/observator/config/server"
	"github.com/srg-bnd/observator/internal/server/router"
)

// Example demonstrates how to create and start an HTTP server
func Example() {
	// Create a router
	router := router.NewRouter(&config.Container{})

	// Create a new server instance
	srv := NewServer(router)

	// Start the server in a separate goroutine
	go func() {
		if err := srv.Start(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Init client
	client := resty.New()

	// Index
	client.R().Get("/")

	// Updates
	client.R().
		SetBody("[{...}]").
		Post("/updates")
}
