// Flags
package main

import (
	"flag"
	"os"
)

var flagHostAddr string

func parseFlags() {
	flag.StringVar(&flagHostAddr, "a", ":8080", "address and port to run server")
	flag.Parse()

	if envHostAddr := os.Getenv("ADDRESS"); envHostAddr != "" {
		flagHostAddr = envHostAddr
	}
}
