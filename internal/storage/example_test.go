package storage

import (
	"context"
	"log"
)

func Example() {
	// Create storage settings
	settings := &Settings{}

	// Initialize storage system
	storage := NewStorage(settings)

	// Example of setting a gauge metric
	ctx := context.Background()
	err := storage.SetGauge(ctx, "example_gauge", 42.5)
	if err != nil {
		log.Fatal(err)
	}

	// Example of retrieving a gauge metric
	_, err = storage.GetGauge(ctx, "example_gauge")
	if err != nil {
		log.Fatal(err)
	}
}
