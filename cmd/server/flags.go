// Flags
package main

import (
	"flag"
	"os"
)

type AppConfigs struct {
	flagHostAddr string
}

var appConfigs = AppConfigs{}

func parseFlags() {
	flag.StringVar(&appConfigs.flagHostAddr, "a", ":8080", "address and port to run server")
	flag.Parse()

	if envHostAddr := os.Getenv("ADDRESS"); envHostAddr != "" {
		appConfigs.flagHostAddr = envHostAddr
	}
}
